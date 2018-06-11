package tests

import (
	"errors"

	db "GoExercises/GORM/models"
)

type UserDataMock struct{}

func (u UserDataMock) GetAllUsersSuccess() ([]db.User, error) {

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

func (u UserDataMock) GetAllUsersFail() ([]db.User, error) {

	return nil, errors.New("Data Can't be fetched from database")
}

func (u UserDataMock) POSTUser() (string, error) {

	return "", errors.New(`{
        "Number": 1062,
        "Message": "Duplicate entry 'Niraj' for key 'api_key'"
    }`)
}
