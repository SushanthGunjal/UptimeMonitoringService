package database

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func CheckUrl(url *Urls) {
	initialCheck := true
	for true {
		if !initialCheck {
			ticker := time.NewTimer(time.Duration(url.Frequency) * time.Second)
			<-ticker.C
		}
		go monitorUrl(url)
		initialCheck = false
	}
}

func monitorUrl(url *Urls) {

	answer, err := httpCalls.MakeHTTPGetRequest(url.CrawlTimeout, url.URL)
	if err != nil {
		IncreaseFailureCount(url) // Request didn't complete within crawl_timeout.
		return
	}
	if answer.StatusCode != http.StatusOK {
		IncreaseFailureCount(url) // Unexpected status-code in response.
	}
	fmt.Println("Completed search on ", url.URL, "at ", time.Now())
}

func IncreaseFailureCount(url *Urls) {
	url.FailureCount++
	if url.FailureCount >= url.FailureThreshold {
		url.Status = "INACTIVE"
		os.Exit(0)
	}
	err := repository.DatabaseSave(url)
	if err != nil {
		log.Fatal(err)
	}
}
