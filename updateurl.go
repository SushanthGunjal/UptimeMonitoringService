package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUrl(c *gin.Context) {
	var url Urls
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&url).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err := repo.databaseSave(&url, id)
	fmt.Println(err)
	c.BindJSON(&url)

	db.Save(&url)
	SetupCloseHandler(url)

	c.JSON(http.StatusOK, url)
}
