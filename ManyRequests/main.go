package main

import (
	"GoExercises/ManyRequests/models"
	"GoExercises/ManyRequests/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/singleflight"
)

//func GetGoID() int32

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	var requestGroup singleflight.Group

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
	router = routes.InitRouters(router, &requestGroup)

	//log.Println("AppName is: ", config.Get("server.appName").(string))

	//router.Run(config.Get("server.httpport").(string))

	fmt.Println("Get All users count at init ", models.GetAllUsersCount)

	router.Run(":8000")

	fmt.Println("Get All users count when program quits ", models.GetAllUsersCount)

}
