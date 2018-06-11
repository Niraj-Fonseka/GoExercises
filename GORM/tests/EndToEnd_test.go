package tests

import (
	controllers "GoExercises/GORM/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGettingAllUsers(t *testing.T) {
	//ts := httptest.NewServer(InitRouters(gin.Default()))
	gin.SetMode(gin.TestMode)

	users, err := UserDataMock{}.GetAllUsersSuccess()

	r := gin.New()
	r.GET("/admin/users", controllers.UserController{users, "", err}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}

func TestGettingAllUsersBadRequest(t *testing.T) {
	//ts := httptest.NewServer(InitRouters(gin.Default()))
	gin.SetMode(gin.TestMode)

	r := gin.New()

	users, err := UserDataMock{}.GetAllUsersFail()

	r.GET("/admin/users", controllers.UserController{users, "", err}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != 500 {
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}

func TestCreatingAUserFail(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	apiKey, err := UserDataMock{}.POSTUser()

	r.POST("/admin/users", controllers.UserController{nil, apiKey, err}.POSTuser)
	req, _ := http.NewRequest("POST", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != 500 {
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}
