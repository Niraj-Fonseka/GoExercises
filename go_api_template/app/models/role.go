package models

import (
	"github.com/jinzhu/gorm"
)

//Role model
type Role struct {
	gorm.Model //Adds fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` in the stuct as well as in the database table, ID is the primary key and auto incremented
	Email      string
	Users      []User `gorm:"many2many:user_roles;"` // for back refrencing ,many to many relationship
}

//TableName Creates table with give name
func (Role) TableName() string {
	return "role"
}
