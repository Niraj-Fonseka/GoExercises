package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine) *gin.Engine {

	//Users
	router.POST("/admin/users", POSTuser)
	router.GET("/admin/users", GETallUsers)

	return router
}
