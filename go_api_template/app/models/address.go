package models

import (
	"github.com/jinzhu/gorm"
)

//Address model
type Address struct {
	gorm.Model //Adds fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` in the stuct as well as in the database table, ID is the primary key and auto incremented
	Address    string
	UserID     uint `gorm:"index"`
}

//TableName Creates table with give name
func (Address) TableName() string {
	return "address"
}
