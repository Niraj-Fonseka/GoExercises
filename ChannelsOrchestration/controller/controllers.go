package controller

import (
	cache "GoExercises/ChannelsOrchestration/cache"

	"github.com/gin-gonic/gin"
)

func HealthRequest(c *gin.Context) {
	cache.CacheDB.Set([]byte("Ping"), []byte(""))
	cache.CacheDB.Set([]byte("Health"), []byte("Health"))
	//check the memory for the

	val, _ := cache.CacheDB.Get([]byte("Health"))
	c.JSON(200, gin.H{
		"message": string(val),
	})
}

func PingTest(c *gin.Context) {
	cache.CacheDB.Set([]byte("Ping"), []byte("Called"))
	cache.CacheDB.Set([]byte("Health"), []byte(""))

	val, _ := cache.CacheDB.Get([]byte("Ping"))

	c.JSON(200, gin.H{
		"message": string(val),
	})
}
