package main

import "github.com/gin-gonic/gin"

func HealthRequest(c *gin.Context) {

	//check the memory for the 
	c.JSON(200, gin.H{
		"message": "Health Ok",
	})
}

func PingTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
