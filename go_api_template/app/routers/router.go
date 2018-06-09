package routers

import (
	"GoExercises/go_api_template/app/controllers"

	jwt "github.com/appleboy/gin-jwt"

	"github.com/gin-gonic/gin"
)

func InitRouters(router *gin.Engine, auth *jwt.GinJWTMiddleware) *gin.Engine {

	router.POST("/login", auth.LoginHandler)
	router.GET("/health", controllers.GetHealth)

	protected := router.Group("/protected")
	protected.Use(auth.MiddlewareFunc())
	{
		protected.GET("/protectedHello", controllers.GetProtectedHealth)
	}

	return router
}
