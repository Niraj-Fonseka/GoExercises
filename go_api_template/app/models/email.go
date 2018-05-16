package models

import (
	"github.com/jinzhu/gorm"
)

//Email model
type Email struct {
	gorm.Model //Adds fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` in the stuct as well as in the database table, ID is the primary key and autoincremented
	Email      string
	UserID     uint `gorm:"index"`
}

//TableName Creates table with give name
func (Email) TableName() string {
	return "email"
}
