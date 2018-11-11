package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// MODELS
type Users []User
type User struct {
	Name string `json"name"`
}

func main() {
	r := SetupRouter()

	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)

	r.Run(":8080")
}

func SetupRouter() *gin.Engine {
	return gin.Default()

}

// ROUTES
func GetUsers(c *gin.Context) {
	repo := UserRepository{}
	ctrl := UserController{}

	ctrl.GetAll(c, repo)
}

func GetUser(c *gin.Context) {
	repo := UserRepository{}
	ctrl := UserController{}

	ctrl.Get(c, repo)
}

// CONTROLLER
type UserController struct{}

func (ctrl UserController) GetAll(c *gin.Context, repository UserRepositoryIterface) {
	c.JSON(200, repository.GetAll())
}

func (ctrl UserController) Get(c *gin.Context, repository UserRepositoryIterface) {

	id := c.Param("id")

	idConv, _ := strconv.Atoi(id)

	c.JSON(200, repository.Get(idConv))
}

// REPOSITORY
type UserRepository struct{}
type UserRepositoryIterface interface {
	GetAll() Users
	Get(id int) User
}

func (r UserRepository) GetAll() Users {
	users := Users{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return users
}

func (r UserRepository) Get(id int) User {
	users := Users{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return users[id-1]
}
