package routers

import (
	controllers "GoExercises/GORM/controllers"
	"GoExercises/GORM/models"

	"github.com/gin-gonic/gin"
)

func GETallUsers(c *gin.Context) {
	data, err := models.UserData{}.GetAllUsers(true)
	ctrl := controllers.UserController{data, err}

	ctrl.GETallUsers(c)
}
