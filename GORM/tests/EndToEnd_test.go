package routers

import (
	controllers "GoExercises/GORM/controllers"
	db "GoExercises/GORM/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGettingAllUsers(t *testing.T) {
	//ts := httptest.NewServer(InitRouters(gin.Default()))
	gin.SetMode(gin.TestMode)

	users, err := db.UserDataMock{}.GetAllUsersSuccess()

	r := gin.New()
	r.GET("/admin/users", controllers.UserController{users, err}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		// If this fails, it's because httprouter needs to be updated to at least f78f58a0db
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}

func TestGettingAllUsersBadRequest(t *testing.T) {
	//ts := httptest.NewServer(InitRouters(gin.Default()))
	gin.SetMode(gin.TestMode)

	r := gin.New()

	users, err := db.UserDataMock{}.GetAllUsersFail()

	r.GET("/admin/users", controllers.UserController{users, err}.GETallUsers)
	req, _ := http.NewRequest("GET", "/admin/users", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != 500 {
		// If this fails, it's because httprouter needs to be updated to at least f78f58a0db
		t.Errorf("Status code should be %v, was %d\n", http.StatusOK, resp.Code)
	}
}
