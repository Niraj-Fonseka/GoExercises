package routers

import (
	controllers "GoExercises/GORM/controllers"
	models "GoExercises/GORM/models"

	"github.com/gin-gonic/gin"
)

type RouterType struct {
	UserData models.UserData
}

func (r RouterType) GETallUsers(c *gin.Context) {
	dbType := r.UserData
	payload, err := controllers.UserController{}.GETallUsers(dbType)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Error",
			"Error":  err.Error,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"users": payload,
		})
		c.Abort()
	}
}
