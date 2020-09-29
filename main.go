package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/SushanthGunjal/UptimeMonitoringService/database"

	uuid "github.com/satori/go.uuid"
)

var db *gorm.DB

type Urls struct {
	ID                uuid.UUID `json:"id" gorm:"primary_key"`
	Url               string    `json:"Url"`
	Crawl_timeout     uint      `json:"Crawl"`
	Frequency         uint      `json:"freq"`
	Failure_threshold uint      `json:"fail"`
}

type extravalue struct {
	Url           string
	Crawl         int
	Freq          int
	Failure       int
	Status        bool
	Failure_count int
}

func main() {
	db = database.Connector()
	r := gin.Default()
	r.POST("/url", CreateUrl)
	r.GET("/url/:id", Geturl)
	r.PATCH("/url/:id", UpdateUrl)
	r.DELETE("url/:id", DeleteUrl)
	r.Run(":8080")

}
