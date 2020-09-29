package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Geturl(c *gin.Context) {
	id := c.Params.ByName("id")
	var url Urls
	if err := db.Where("id = ?", id).First(&url).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		err := repo.databaseGet(&url)
		fmt.Println(err)
		SetupCloseHandler(url)
		c.JSON(http.StatusOK, url)
		checkurl(url)
	}
}
