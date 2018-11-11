package main

import (
	cache "GoExercises/ChannelsOrchestration/cache"
	controller "GoExercises/ChannelsOrchestration/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//Opening ledis database
	cache.InitLedisDB()
	r.GET("/ping", controller.PingTest)
	r.GET("/health", controller.HealthRequest)
	r.Run() // listen and serve on 0.0.0.0:8080
}
