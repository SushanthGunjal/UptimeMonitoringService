package database

import (
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCheckUrls(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := NewMockRepositoryController(ctrl)
	mockHttp := NewMockHttpController(ctrl)

	mockRepo.EXPECT().DatabaseSave(gomock.Any()).Return(nil).MaxTimes(10)
	mockRepo.EXPECT().DatabaseGet(gomock.Any()).Return(nil).MaxTimes(10)
	SetRepoController(mockRepo)

	// Case 1: status_code == StatusOK
	mockHttp.EXPECT().MakeHTTPGetRequest(gomock.Any(), "https://example1.com").Return(
		&http.Response{StatusCode: http.StatusOK}, nil)

	// Case 2: status_code != StatusOK
	mockHttp.EXPECT().MakeHTTPGetRequest(gomock.Any(), "https://example2.com").Return(
		&http.Response{StatusCode: http.StatusConflict}, nil)

	SetHTTPController(mockHttp)
	httpCalls = GetHTTPController()
	repository = GetRepoController()

	url1 := Urls{
		ID:               "Case1",
		FailureCount:     0,
		FailureThreshold: 5,
		Status:           "ACTIVE",
		CrawlTimeout:     20,
		URL:              "https://example1.com",
	}
	monitorUrl(&url1)
	assert.Equal(t, url1.FailureCount, 0)
	assert.Equal(t, url1.Status, "ACTIVE")

	url2 := Urls{
		ID:               "Case2",
		FailureCount:     0,
		FailureThreshold: 5,
		Status:           "ACTIVE",
		CrawlTimeout:     20,
		URL:              "https://example2.com",
	}
	monitorUrl(&url2)
	assert.Equal(t, url2.FailureCount, 1)
	assert.Equal(t, url2.Status, "ACTIVE")

}
