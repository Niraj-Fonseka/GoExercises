package models

import "errors"

type UserDataMock struct{}

func (u UserDataMock) GetAllUsersSuccess() ([]User, error) {

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

func (u UserDataMock) GetAllUsersFail() ([]User, error) {

	return nil, errors.New("Data Can't be fetched from database")
}
