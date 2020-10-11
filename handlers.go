package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SushanthGunjal/UptimeMonitoringService/database"

	"github.com/gin-gonic/gin"
)

//Adds an URL to the database.
func CreateUrl(c *gin.Context) {
	var postRequest database.PostRequest
	err := c.BindJSON(&postRequest)
	if err != nil {
		log.Fatal(err)
		return
	}
	url, err := database.AddintoDB(postRequest)
	if err != nil {
		log.Fatal(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":                url.ID,
			"url":               url.URL,
			"crawl_timeout":     url.CrawlTimeout,
			"frequency":         url.Frequency,
			"failure_threshold": url.FailureThreshold,
			"status":            url.Status,
			"failure_count":     url.FailureCount,
		})

	}
	fmt.Println(url)
}

//Activates an Url in the database
func ActivateUrl(c *gin.Context) {
	id := c.Param("id")
	if url, flag := database.GetUrlbyId(id); flag {
		database.ActivateUrlInDB(url)
		c.JSON(http.StatusOK, gin.H{
			"id":     url.ID,
			"status": url.Status,
		})
	} else {
		fmt.Println("invalid id ")
	}
}

//Deactivates an Url in the database
func DeactivateUrl(c *gin.Context) {
	id := c.Param("id")
	if url, flag := database.GetUrlbyId(id); flag {
		database.DeActivateUrlInDB(url)
		c.JSON(http.StatusOK, gin.H{
			"id":     url.ID,
			"status": url.Status,
		})
	} else {
		fmt.Println("invalid id ")
	}

}

//Searches Url from the database
func Geturl(c *gin.Context) {
	id := c.Param("id")
	//fmt.Println(id)
	if url, flag := database.GetUrlbyId(id); flag {
		go database.CheckUrl(url)
		c.JSON(http.StatusOK, gin.H{
			"id":                url.ID,
			"url":               url.URL,
			"crawl_timeout":     url.CrawlTimeout,
			"frequency":         url.Frequency,
			"failure_threshold": url.FailureThreshold,
			"status":            url.Status,
			"failure_count":     url.FailureCount,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}
}

//Updates an Url in the database
func UpdateUrl(c *gin.Context) {
	var patchRequest database.PatchRequest
	err := c.BindJSON(&patchRequest)

	if err != nil {
		log.Fatal(err)
	}

	id := c.Param("id")
	if url, flag := database.UpdateURLById(id, patchRequest); flag {
		c.JSON(http.StatusOK, gin.H{
			"id":                url.ID,
			"url":               url.URL,
			"crawl_timeout":     url.CrawlTimeout,
			"frequency":         url.Frequency,
			"failure_threshold": url.FailureThreshold,
			"status":            url.Status,
			"failure_count":     url.FailureCount,
		})

	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})

	}
}

//Deletes an Url from the database
func DeleteUrl(c *gin.Context) {
	id := c.Param("id")
	flag := database.DeleteUrlById(id)
	if !flag {
		c.JSON(http.StatusNotFound, gin.H{
			"id": "invalid",
		})
	}

}
