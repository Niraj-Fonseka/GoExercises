package routers

import (
	"GoExercises/GoTest/DbApp/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	//Users
	router.POST("/admin/users", controllers.POSTuser)
	router.GET("/admin/users", controllers.GETallUsers)
	router.GET("/address", controllers.GETallAddresses)
	router.POST("/address", controllers.POSTaddress)

	return router
}
