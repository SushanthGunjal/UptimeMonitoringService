package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func checkurl(url Urls) {
	count := 0
	status := true
	channel := make(chan string)
	link := url.Url
	crawl := int(url.Crawl_timeout)
	failure_threshold := int(url.Failure_threshold)
	fmt.Println(link)
	go checklink(link, channel, &count, failure_threshold, &status, crawl)
	for _ = range channel {
		go func(checkedlink string) {
			if status {
				frequency := time.Duration(url.Frequency)
				time.Sleep(frequency * time.Second)
				checklink(checkedlink, channel, &count, failure_threshold, &status, crawl)
			}

		}(link)
	}
}
func checklink(link string, channel chan string, count *int, failure_threshold int, status *bool, crawl int) {

	dt := time.Now()
	fmt.Println("started new search on ", link, "at ", dt)
	timeout := time.Duration(time.Duration(crawl) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	response, err := client.Get(link)
	if err != nil || response.StatusCode != http.StatusOK {
		*count++
		if *count >= failure_threshold {
			*status = false
			fmt.Println("Your lifetime is finished", link)

		}
		fmt.Println(err)
		fmt.Println("The failure count of ", link, "is", *count)
		fmt.Println("failed to search", link)
		fmt.Println()

		channel <- link
		return
	}
	fmt.Println("Succesful Search", link)
	fmt.Println()
	channel <- link
	return
}

func IsUrl(str string) (bool, string) {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != "", u.Scheme + "://" + u.Host + u.Path
}
