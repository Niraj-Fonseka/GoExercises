package routers

import (
	"go_exercises/GORM/controllers"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	//Users
	router.POST("/admin/users", controllers.POSTuser)
	router.GET("/admin/users", controllers.GETallUsers)

	return router
}
