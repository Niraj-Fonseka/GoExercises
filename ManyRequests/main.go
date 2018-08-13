package main

import (
	"GoExercises/ManyRequests/models"
	"GoExercises/ManyRequests/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//func GetGoID() int32

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	//Opening the database
	models.OpenDB(os.Getenv("databaseConn"), false) // Pass false for not logging dababase queries
	defer models.CloseDB(models.DB)

	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	log.SetOutput(gin.DefaultWriter)             // logs gin
	log.SetFlags(log.LstdFlags | log.Lshortfile) // logs with timestamp and file name and line number

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	//auth := auth.InitAuthentication()
	router = routes.InitRouters(router)

	//log.Println("AppName is: ", config.Get("server.appName").(string))

	//router.Run(config.Get("server.httpport").(string))
	router.Run(":8000")
}
