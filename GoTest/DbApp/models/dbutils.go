package models

//ALTER TABLE `service` ADD `id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/jinzhu/gorm"
)

//DB object to be used
var DB *gorm.DB

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
	err = DB.AutoMigrate(&User{}, &Address{}).Error

	DB.Model(&Address{}).AddForeignKey("api_key", "user(api_key)", "CASCADE", "CASCADE")

	if err != nil {
		log.Panicln("Error in Migrating tables", err.Error())

		return err
	}

	// Set conection pool
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return nil
}

func TestDBInit() *gorm.DB {
	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
	if err != nil {
		fmt.Println("db err: ", err)
	}
	test_db.DB().SetMaxIdleConns(3)
	test_db.LogMode(true)
	DB = test_db
	return DB
}

// Delete the database after running testing cases.
func TestDBFree(test_db *gorm.DB) error {
	test_db.Close()
	err := os.Remove("./../gorm_test.db")
	return err
}

//CloseDB closes db connection
func CloseDB(db *gorm.DB) {

	db.Close()
}
