package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/SushanthGunjal/UptimeMonitoringService/database"
)

func main() {
	database.Connector()
	r := gin.Default()
	r.POST("/url", CreateUrl)
	r.GET("/url/:id", Geturl)
	r.PATCH("/url/:id", UpdateUrl)
	r.DELETE("url/:id", DeleteUrl)
	r.POST("url/:id/activate", ActivateUrl)
	r.POST("url/:id/deactivate", DeactivateUrl)
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
