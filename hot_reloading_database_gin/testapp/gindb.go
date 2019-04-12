package main

import (
	"fmt"
	"log"
	"net/http"
	// "strconv"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func init() {
	fmt.Println("Initializing App")
}

var DB *gorm.DB

func main() {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	// we should eventually specify our origins for better security
	corsConfig.AllowAllOrigins = true

	// Register database
	log.Println("Registering database..")
	OpenDB("root:password@tcp(localhost:3333)/dbname?charset=utf8&parseTime=True&loc=Local", false) // Pass false for not logging dababase queries
	defer CloseDB(DB)

	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	log.SetOutput(gin.DefaultWriter)             // logs gin
	log.SetFlags(log.LstdFlags | log.Lshortfile) // logs with timestamp and file name and line number

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router = InitRouters(router)

	//log.Println("AppName is: ", config.Get("server.appName").(string))

	router.Run(":8080")

}

func InitRouters(router *gin.Engine) *gin.Engine {

	router.GET("/health", GetHealth)

	return router
}

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Health Status : OK !")
}

//OpenDB connects database with settings
func OpenDB(databaseConn string, mode bool) (err error) {

	log.Println("Opening database ..", databaseConn)

	DB, err = gorm.Open("mysql", databaseConn)
	if err != nil {

		log.Fatal("Error in database connection ", err.Error())
		return err
	}

	// Enable Logger, show detailed log
	DB.LogMode(mode)

	// updating schema , add databse models here to autogenerate schema
	//err = DB.AutoMigrate(&User{}).Error

	if err != nil {
		log.Panicln("Error in Migrating tables", err.Error())

		return err
	}

	// Set conection pool
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return nil
}

//CloseDB closes db connection
func CloseDB(db *gorm.DB) {

	db.Close()
}
