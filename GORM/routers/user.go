package routers

import (
	controllers "GoExercises/GORM/controllers"
	models "GoExercises/GORM/models"

	"github.com/gin-gonic/gin"
)

func GETallUsers(c *gin.Context) {
	data := models.UserData{}
	controller := controllers.UserController{}

	controller.GETallUsers(c, data)
}
