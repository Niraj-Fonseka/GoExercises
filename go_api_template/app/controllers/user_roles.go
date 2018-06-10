package controllers

import (
	db "GoExercises/go_api_template/app/models"

	"github.com/gin-gonic/gin"
)

func GETallUserRoles(c *gin.Context) {
	allRoles, err := db.GetAllUserRoles(true)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"user_roles": allRoles,
		})
	}
}

func POSTuserRole(c *gin.Context) {
	var createRole db.UserRoles
	err := c.BindJSON(&createRole)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Possible Invalid JSON Body",
			"err":    err,
		})
		c.Abort()
	}

	api_key, err := db.CreateUserRole(&createRole)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status":  "Role Created",
			"api_key": api_key,
		})
	}
}
