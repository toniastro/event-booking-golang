package models

import (
	"github.com/jinzhu/gorm"
	"os"
	"github.com/joho/godotenv"
	"github.com/ichtrojan/thoth"
	"log"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB //database

func init() {
	logger, _ := thoth.Init("log")

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")


	db_details := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",username,password, dbHost, dbName) //Build connection string
	fmt.Println(db_details)

	conn, err := gorm.Open("mysql", db_details)

	if err != nil {
		logger.Log(err)
		log.Fatal(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Details{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}