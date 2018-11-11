package models

import (
	"testing"

	_ "github.com/gin-gonic/gin"
)

type UserDataMock struct{}

func (u UserDataMock) GetAllUsersTest() ([]User, error) {

	users := []User{
		User{
			Api_key:    "username 1",
			Api_secret: "secret1",
		},
		User{
			Api_key:    "username 2",
			Api_secret: "secret2",
		},
		User{
			Api_key:    "username 3",
			Api_secret: "secret3",
		},
		User{
			Api_key:    "username 4",
			Api_secret: "secret4",
		},
	}

	return users, nil
}

// TESTING REPOSITORY FUNCTIONS
func TestRepoGetAll(t *testing.T) {

	data := UserData{}

	users, _ := data.GetAllUsers(true)

	amountUsers := len(users)

	if amountUsers != 2 {
		t.Errorf("Esperado %d, recebido %d", 2, amountUsers)
	}
}
