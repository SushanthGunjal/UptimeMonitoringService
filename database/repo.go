package database

import (
	"errors"
	"net/http"
	"time"

	"github.com/gojektech/heimdall/httpclient"
)

type RepositoryController interface {
	DatabaseCreate(url *Urls) error
	DatabaseGet(url *Urls) error
	DatabaseSave(url *Urls) error
	DatabaseDelete(url *Urls) error
}
type HttpController interface {
	MakeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error)
}

type MonitorRepo struct{}

type MonitorHttp struct{}

var repository RepositoryController

var httpCalls HttpController

func SetRepoController(repoType RepositoryController) {
	repository = repoType
}

func GetRepoController() RepositoryController {
	return repository
}

func (rp *MonitorRepo) DatabaseCreate(url *Urls) error {
	return DB.Create(&url).Error
}

func (rp *MonitorRepo) DatabaseGet(url *Urls) error {
	answer := DB.First(&url)
	if answer.RowsAffected == 0 {
		return errors.New("Url not found")
	}
	return answer.Error
}

func (rp *MonitorRepo) DatabaseSave(url *Urls) error {
	return DB.Save(&url).Error
}

func (rp *MonitorRepo) DatabaseDelete(url *Urls) error {
	answer := DB.Delete(&url)
	if answer.RowsAffected == 0 {
		return errors.New("Url not found")
	}
	return answer.Error
}

func SetHTTPController(hType HttpController) {
	httpCalls = hType
}

func GetHTTPController() HttpController {
	return httpCalls
}

func (mh *MonitorHttp) MakeHTTPGetRequest(crawlTimeout int, url string) (*http.Response, error) {

	timeout := time.Duration(crawlTimeout) * time.Second
	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	return client.Get(url, nil)
}
