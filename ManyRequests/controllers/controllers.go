package controllers

import (
	db "GoExercises/ManyRequests/models"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

type Controller struct {
	RGroup *singleflight.Group
}

func GetHealth(c *gin.Context) {

	var requestGroup singleflight.Group

	v, err, shared := requestGroup.Do("health", func() (interface{}, error) {

		//fmt.Println(GetGoID())

		go func() {
			time.Sleep(30 * time.Second)
			log.Println("Deleting \"health\" key")
			requestGroup.Forget("health")
		}()

		return db.GetHealth()
	})

	status := v.(string)
	fmt.Println(err)
	fmt.Println("Value :  ", v)
	log.Printf("/getallusers handler requst: status %q, shared result %t", status, shared)
	log.Printf("Value : ", v)
	fmt.Printf("Status is : %s \n", status)

	c.JSON(200, v)

}

func (ctrl Controller) GetAllUsers(c *gin.Context) {

	URL := c.Request.URL.String()
	fmt.Printf("URL : %s \n", URL)
	v, err, _ := ctrl.RGroup.Do(URL, func() (interface{}, error) {

		//fmt.Println(GetGoID())

		go func() {
			time.Sleep(30 * time.Second)
			log.Println("Deleting key ", URL)
			ctrl.RGroup.Forget(URL)
		}()

		users, err := db.GetAllUsers(false)

		value, err := json.Marshal(users)

		return string(value), err
	})

	//status := v.(string)
	//fmt.Println(err)
	// fmt.Println("Value :  ", v)
	// log.Printf("/getallusers handler requst: status %q, shared result %t", status, shared)
	// log.Printf("Value :")
	// fmt.Printf("Status is : %s \n", status)

	//val, err := db.GetAllUsers(false)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "get failed",
			"err":    err,
		})
	} else {
		c.JSON(200, gin.H{
			"users": v,
		})
	}

}

func AddUser(c *gin.Context) {
	var createUser db.User
	err := c.BindJSON(&createUser)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Possible Invalid JSON Body",
			"err":    err,
		})
		c.Abort()
	}

	userName, err := db.CreateUser(&createUser)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status":    "user created",
			"user_name": userName,
		})
	}
}

func GetUser(c *gin.Context) {
	userName := c.Param("username")
	userName = string(userName)

	users, err := db.GetUserByID(userName, false)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "get failed",
			"err":    err,
		})
	} else {
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}
