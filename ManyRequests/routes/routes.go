package routes

import (
	"GoExercises/ManyRequests/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	router.GET("/health", controllers.GetHealth)

	router.POST("/user", controllers.AddUser)
	router.GET("/users", controllers.GetAllUsers)
	router.GET("/user/:username", controllers.GetUser)

	return router
}
