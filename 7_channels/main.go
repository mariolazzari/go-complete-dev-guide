package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://bing.com",
		"https://mariolazzari.it",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func() {
			time.Sleep(time.Second * 5)
			checkLink(l, ch)
		}()
	}
}

func checkLink(link string, ch chan string) {
	time.Sleep(time.Second * 5)
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
