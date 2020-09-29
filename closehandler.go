package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SetupCloseHandler(url Urls) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		//DeleteFiles()
		fmt.Println("id: ", url.ID)
		fmt.Println("url: ", url.Url)
		fmt.Println("crawl_timeout: ", url.Crawl_timeout)
		fmt.Println("frequency: ", url.Frequency)
		fmt.Println("failure_threshold: ", url.Failure_threshold)
		os.Exit(0)
	}()
}
