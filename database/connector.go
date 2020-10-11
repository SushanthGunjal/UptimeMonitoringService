package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

//Connects to the database.
func Connector() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	DB_User := os.Getenv("DB_USER")
	DB_Password := os.Getenv("DB_PASSWORD")
	DB_Port := os.Getenv("DB_PORT")
	DB_Name := os.Getenv("DB_NAME")
	dsn := DB_User + ":" + DB_Password + "@" + "(" + DB_Port + ")" + "/" + DB_Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection succesfull", DB)
	}
	err = DB.AutoMigrate(&Urls{}) // Makes the table of structure Urls
	if err != nil {
		os.Exit(1)
	}
	//var url Urls

}
