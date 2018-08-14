package routes

import (
	"GoExercises/ManyRequests/controllers"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

func InitRouters(router *gin.Engine, requestGroup *singleflight.Group) *gin.Engine {

	router.GET("/health", controllers.GetHealth)

	router.POST("/user", controllers.AddUser)
	router.GET("/users", controllers.Controller{RGroup: requestGroup}.GetAllUsers)
	router.GET("/user/:username", controllers.GetUser)

	return router
}
