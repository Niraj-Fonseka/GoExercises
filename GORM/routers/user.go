package routers

import (
	controllers "GoExercises/GORM/controllers"
	"GoExercises/GORM/models"

	"github.com/gin-gonic/gin"
)

func GETallUsers(c *gin.Context) {
	data, err := models.UserData{}.GetAllUsers(true)
	ctrl := controllers.UserController{data, "", err}

	ctrl.GETallUsers(c)
}

func POSTuser(c *gin.Context) {
	var createUser models.User
	err := c.BindJSON(&createUser)

	apiKey, err := models.UserData{}.CreateUser(&createUser)
	ctrl := controllers.UserController{nil, apiKey, err}
	ctrl.POSTuser(c)
}
