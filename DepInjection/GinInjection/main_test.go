package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type UserRepositoryMock struct{}

func (r UserRepositoryMock) GetAll() Users {
	users := Users{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return users
}

func (r UserRepositoryMock) Get(id int) User {
	users := Users{
		{Name: "Wilson"},
		{Name: "Panda"},
	}

	return users[id-1]
}

// TESTING REPOSITORY FUNCTIONS
func TestRepoGetAll(t *testing.T) {

	userRepo := UserRepository{}

	amountUsers := len(userRepo.GetAll())

	if amountUsers != 2 {
		t.Errorf("Esperado %d, recebido %d", 2, amountUsers)
	}
}

func TestRepoGet(t *testing.T) {

	expectedUser := struct {
		Name string
	}{
		"Wilson",
	}

	userRepo := UserRepository{}

	user := userRepo.Get(1)

	if user.Name != expectedUser.Name {
		t.Errorf("Esperado %s, recebido %s", expectedUser.Name, user.Name)
	}
}

func TestPingRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
