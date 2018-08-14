package models

import (
	"fmt"
	"sync"
	"time"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Email    string `gorm:"not null;unique;size:100"" json:"email"`
	UserName string `json:"user_name"`
	Age      int    `json:"age"`
}

//TableName Creates table with give name
func (User) TableName() string {
	return "users"
}

var mu sync.Mutex

func CreateUser(user *User) (un string, err error) {

	if err = DB.Create(user).Error; err == nil {
		return user.UserName, nil
	}
	return "", err
}

//ModifyUser modifies data
func ModifyUser(user *User) (err error) {

	if err = DB.Save(user).Error; err == nil {
		return nil
	}
	return err
}

//GetUsertByID gets User by id
//keep autoPreload true if needs to fetch all related data else keep false
func GetUserByID(un string, autoPreload bool) (user *User, err error) {
	user = &User{}
	if autoPreload == true {
		if err = DB.Set("gorm:auto_preload", true).First(user, un).Error; err != nil {

			return nil, err
		}
	} else {
		if err = DB.First(user, un).Error; err != nil {

			return nil, err
		}
	}
	return user, nil
}

func GetAllUsers(autoPreload bool) (users []User, err error) {
	users = []User{}

	mu.Lock()
	GetAllUsersCount++
	mu.Unlock()
	fmt.Println("Doing the DB Call for getting all users")
	time.Sleep(5 * time.Second)

	if autoPreload == true {
		err = DB.Set("gorm:auto_preload", true).Find(&users).Error
	} else {
		err = DB.Find(&users).Error
	}

	if err == nil {
		return users, nil
	}
	return nil, err
}
