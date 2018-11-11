package controllers

import (
	db "GoExercises/GoTest/DbApp/models"

	"github.com/gin-gonic/gin"
)

func GETallAddresses(c *gin.Context) {
	addrs, err := db.GetAllAddresses(true)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "get failed",
			"err":    err,
		})
	} else {
		c.JSON(200, gin.H{
			"addresses": addrs,
		})
	}
}

func POSTaddress(c *gin.Context) {
	var createAddress db.Address
	err := c.BindJSON(&createAddress)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Possible Invalid JSON Body",
			"err":    err,
		})
		c.Abort()
	}

	addressStreet, err := db.CreateAddress(&createAddress)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status": "Address Created",
			"street": addressStreet,
		})
	}
}
