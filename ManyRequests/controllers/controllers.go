package controllers

import (
	db "GoExercises/ManyRequests/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Health Status : OK !")
}

func GetAllUsers(c *gin.Context) {
	users, err := db.GetAllUsers(false)
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
