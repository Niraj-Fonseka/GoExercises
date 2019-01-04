package main

import (
	"fmt"
	"log"
	"net/http"
	// "strconv"
	models "./models"

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
	corsConfig.AllowAllOrigins = true

	// Register database
	log.Println("Registering database..")

	//OpenDB(os.Getenv("DBCON"), false) // Pass false for not logging dababase queries

	gin.DisableConsoleColor()

	log.SetOutput(gin.DefaultWriter)             // logs gin
	log.SetFlags(log.LstdFlags | log.Lshortfile) // logs with timestamp and file name and line number

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router = InitRouters(router)

	router.Run(":8080")

}

func InitRouters(router *gin.Engine) *gin.Engine {

	router.GET("/health", GetHealth)

	protected := router.Group("/protected")
	{
		protected.GET("/protectedHealth", GetProtectedHealth)
	}

	router.GET("/users", GETallUsers)

	userGroup := router.Group("/user")
	userGroup.POST("/", POSTuser)
	userGroup.GET("/:userID", GETUserByApiKey)

	return router
}

//Controller

func GETallUsers(c *gin.Context) {
	users, err := models.GetAllUsers(true)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}

func GETUserByApiKey(c *gin.Context) {
	userID := c.Param("userID")
	user, err := models.GetUserByFeild("api_key", userID)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "GET failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func POSTuser(c *gin.Context) {
	var createUser models.User
	err := c.BindJSON(&createUser)

	if err != nil {
		c.JSON(500, gin.H{
			"status": "Possible Invalid JSON Body",
			"err":    err,
		})
		c.Abort()
	}

	api_key, err := models.CreateUser(&createUser)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "post failed",
			"err":    err,
		})
		c.Abort()
	} else {
		c.JSON(200, gin.H{
			"status":  "API key and secret created",
			"api_key": api_key,
		})
	}
}

func GetHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Health Status : OK !")
}

func GetProtectedHealth(c *gin.Context) {
	c.JSON(http.StatusOK, "Protected Health Status : OK !")
}

//Database
