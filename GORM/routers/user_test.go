package routers

import (
	controllers "GoExercises/GORM/controllers"
	db "GoExercises/GORM/models"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type RouterTypeMock struct {
	UserData UserDataMock
}

type UserDataMock struct {
	Users []db.User
}

type UserDataMockInterface interface {
	GetAllUsers(autoPreload bool) ([]db.User, error)
}

func (u UserDataMock) GetAllUsers() ([]db.User, error) {

	users := []db.User{
		db.User{
			Api_key:    "username 1",
			Api_secret: "secret1",
		},
		db.User{
			Api_key:    "username 2",
			Api_secret: "secret2",
		},
		db.User{
			Api_key:    "username 3",
			Api_secret: "secret3",
		},
		db.User{
			Api_key:    "username 4",
			Api_secret: "secret4",
		},
	}

	return users, nil
}

var users = []db.User{
	db.User{
		Api_key:    "username 1",
		Api_secret: "secret1",
	},
	db.User{
		Api_key:    "username 2",
		Api_secret: "secret2",
	},
	db.User{
		Api_key:    "username 3",
		Api_secret: "secret3",
	},
	db.User{
		Api_key:    "username 4",
		Api_secret: "secret4",
	},
}

func TestGettingAllUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	UserDataMockVar := UserDataMock{}
	r.GET("/admin/users", RouterType{UserDataMock{users}}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	fmt.Println(resp.Body.String())
	fmt.Println(resp.Code)

	if resp.Code != http.StatusOK {
		// If this fails, it's because httprouter needs to be updated to at least f78f58a0db
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}

func TestGettingAllUsersBadRequest(t *testing.T) {
	//ts := httptest.NewServer(InitRouters(gin.Default()))
	gin.SetMode(gin.TestMode)

	r := gin.New()
	err := errors.New("Bad Request")
	r.GET("/admin/users", controllers.UserController{nil, err}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	fmt.Println(resp.Body.String())
	fmt.Println(resp.Code)

	if resp.Code != 500 {
		// If this fails, it's because httprouter needs to be updated to at least f78f58a0db
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}
