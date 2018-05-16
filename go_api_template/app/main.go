package main

import (
	"fmt"
	"log"
	"os"
	"valet-crud-api_go/app/auth"
	"valet-crud-api_go/app/models"
	"valet-crud-api_go/app/routers"
	// "strconv"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	fmt.Println("Initializing App")
}

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	// we should eventually specify our origins for better security
	corsConfig.AllowAllOrigins = true

	// Register database
	log.Println("Registering database..")

	// Add Database connection conn for MySql = "username:password@/dbname?charset=utf8&parseTime=True&loc=Local"

	// Add Database connection conn for CloundMySql = "username:password123@(cloudsqlinstance)/dbname?charset=utf8&parseTime=True&loc=UTC"

	models.OpenDB(os.Getenv("dbconn"), false) // Pass false for not logging dababase queries
	defer models.CloseDB(models.DB)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// // Logging to a file.
	// f, err := os.Create(config.Get("logLocation").(string))

	// if err != nil {
	// 	log.Panicln("error is ", err.Error())
	// }

	// Use the following code if you need to write the logs to file and console at the same time.
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	log.SetOutput(gin.DefaultWriter)             // logs gin
	log.SetFlags(log.LstdFlags | log.Lshortfile) // logs with timestamp and file name and line number

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	auth := auth.InitAuthentication()
	router = routers.InitRouters(router, auth)

	//log.Println("AppName is: ", config.Get("server.appName").(string))

	router.Run(":8080")

}
