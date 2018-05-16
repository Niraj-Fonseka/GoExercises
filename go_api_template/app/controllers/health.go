package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Health Status : OK !")
}

func GetProtectedHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Protected Health Status : OK !")
}
