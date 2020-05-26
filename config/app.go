package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*var (
	db *gorm.DB
)

// Connect cheks the connection with the database
func Connect() {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := " "
	dbName := "db_feedback"
	d, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3308)/"+dbName)

	// if there is an error opening the connection, handle it
	if err == nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}*/

const (
	DB_USER = "root"
	DB_PASS = ""
	DB_HOST = "127.0.0.1"
	DB_PORT = 3308
	DB_NAME = "db_feedback"
)

// Connect to connect the db
func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
