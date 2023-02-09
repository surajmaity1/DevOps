package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wait_grp sync.WaitGroup

func main() {
	/*
		go caller("Hey")
		caller("Hi")
	*/

	websites := []string{
		"https://google.com",
		"https://discord.com",
		"https://slack.com",
		"https://telegram.com",
		"https://reddit.com",
	}

	for _, website := range websites {
		go getStatusCode(website)
		wait_grp.Add(1)
	}

	wait_grp.Wait()
}

func caller(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	}
}
func getStatusCode(endpoint string) {
	defer wait_grp.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("###### ERROR #######")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}
