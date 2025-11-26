package main

import (
	"fmt"
	"net/http"
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
		go checkLink(l, ch)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- link
		return
	}

	ch <- link
}
