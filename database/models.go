package database

type Urls struct {
	ID               string `gorm:"primaryKey"`
	URL              string
	CrawlTimeout     int
	Frequency        int
	FailureThreshold int
	Status           string
	FailureCount     int
}

type PostRequest struct {
	URL              string `json:"url"`
	CrawlTimeout     int    `json:"crawl_timeout"`
	Frequency        int    `json:"frequency"`
	FailureThreshold int    `json:"failure_threshold"`
}

type PatchRequest struct {
	Frequency        int `json:"frequency"`
	FailureThreshold int `json:"failure_threshold"`
	CrawlTimeout     int `json:"crawl_timeout"`
}
