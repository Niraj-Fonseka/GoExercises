package controllers

import (
	db "GoExercises/go_api_template/app/models"

	"github.com/gin-gonic/gin"
)

func GETallUsers(c *gin.Context) {
	users, err := db.GetAllUsers(true)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}

func GETUserByApiKey(c *gin.Context) {
	userID := c.Param("userID")
	user, err := db.GetUserByFeild("api_key", userID)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "GET failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

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
