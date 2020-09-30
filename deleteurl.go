package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//Deletes an URL from the database.
func DeleteUrl(c *gin.Context) {
	id := c.Params.ByName("id")
	var url Urls
	err := repo.databaseDelete(&url, id)
	fmt.Println(err)
	data := db.Where("id = ?", id).Delete(&url)
	fmt.Println(data)
	c.JSON(200, gin.H{"id " + id: "deleted"})
}
