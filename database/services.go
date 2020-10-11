package database

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

func DeleteUrlById(id string) bool {
	url, flag := GetUrlbyId(id)
	if !flag {
		return false
	}
	err := repository.DatabaseDelete(url)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func AddintoDB(req PostRequest) (Urls, error) {
	id := uuid.New().String()
	urlInfo := Urls{
		ID:               id,
		URL:              req.URL,
		CrawlTimeout:     req.CrawlTimeout,
		Frequency:        req.Frequency,
		FailureThreshold: req.FailureThreshold,
		Status:           "ACTIVE",
		FailureCount:     0,
	}
	repository.DatabaseCreate(&urlInfo)
	return urlInfo, nil
}

func GetUrlbyId(id string) (*Urls, bool) {
	url := Urls{
		ID: id,
	}
	err := repository.DatabaseGet(&url)
	if err != nil {
		log.Fatal(err)
	}
	if url.URL == "" {
		return &Urls{}, false
	}
	fmt.Println(url)
	return &url, true
}

func UpdateURLById(id string, patchrequest PatchRequest) (*Urls, bool) {
	url, flag := GetUrlbyId(id)
	if !flag {
		return &Urls{}, false
	}
	url.Frequency = patchrequest.Frequency
	url.FailureThreshold = patchrequest.FailureThreshold
	url.CrawlTimeout = patchrequest.CrawlTimeout
	url.Status = "ACTIVE"
	url.FailureCount = 0
	err := repository.DatabaseSave(url)
	if err != nil {
		log.Fatal(err)
	}
	return url, true
}

func ActivateUrlInDB(url *Urls) (*Urls, bool) {
	url.Status = "ACTIVE"
	err := repository.DatabaseSave(url)
	if err != nil {
		log.Fatal(err)
	}
	return url, true
}

func DeActivateUrlInDB(url *Urls) (*Urls, bool) {
	url.Status = "INACTIVE"
	err := repository.DatabaseSave(url)
	if err != nil {
		log.Fatal(err)
	}
	return url, true
}
