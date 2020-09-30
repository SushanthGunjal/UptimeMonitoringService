package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connects to the database.
func Connector() *gorm.DB {
	dsn := "root:mysql@tcp(docker.for.mac.localhost:3306)/checkurl?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connection succesfull", db)
	return db
}
