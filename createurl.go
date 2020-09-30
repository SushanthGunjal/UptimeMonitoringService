package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

//Adds an URL to the database.
func CreateUrl(c *gin.Context) {
	var url Urls
	url.ID = uuid.NewV4()
	c.BindJSON(&url)
	err := repo.databaseCreate(&url)
	fmt.Println(err)
	db.Create(&url)
	SetupCloseHandler(url)
	c.JSON(http.StatusOK, url)
}
