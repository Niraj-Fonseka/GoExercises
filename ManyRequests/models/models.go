package models

//ALTER TABLE `service` ADD `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

//DB object to be used
var DB *gorm.DB
var GetAllUsersCount int64

//OpenDB connects database with settings
func OpenDB(databaseConn string, mode bool) (err error) {

	log.Println("Opening database ..", databaseConn)

	DB, err = gorm.Open("mysql", databaseConn)
	if err != nil {

		log.Fatal("Error in database connection ", err.Error())
		return err
	}

	// Enable Logger, show detailed log
	DB.LogMode(mode)

	// updating schema , add databse models here to autogenerate schema
	err = DB.AutoMigrate(&User{}).Error

	if err != nil {

		log.Panicln("Error in Migrating tables", err.Error())

		return err
	}

	// Set conection pool
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return nil
}

//CloseDB closes db connection
func CloseDB(db *gorm.DB) {

	db.Close()
}
