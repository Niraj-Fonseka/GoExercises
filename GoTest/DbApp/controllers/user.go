package controllers

import (
	db "GoExercises/GoTest/DbApp/models"

	"github.com/gin-gonic/gin"
)

func POSTuser(c *gin.Context) {
	var createUser db.User
	err := c.BindJSON(&createUser)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Possible Invalid JSON Body",
			"err":    err,
		})
		c.Abort()
	}

	api_key, err := db.CreateUser(&createUser)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status":  "API key and secret created",
			"api_key": api_key,
		})
	}
}

func GETallUsers(c *gin.Context) {
	users, err := db.GetAllUsers(true)
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

// func DELETEUser(c *gin.Context) {

// 	var deleteUser db.User
// 	err := c.BindJSON(&deleteUser)
// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"status": "Possible Invalid JSON Body",
// 			"err":    err,
// 		})
// 		c.Abort()
// 	}

// 	err = db.DeleteUser(deleteUser, true)

// 	if err != nil {
// 		c.JSON(500, gin.H{
// 			"status": "delete failed",
// 			"err":    err,
// 		})
// 		c.Abort()
// 	} else {
// 		c.JSON(200, gin.H{
// 			"status":  "User Deleted Successfully",
// 			"api_key": deleteUser.Api_key,
// 		})
// 	}
// }
