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
		// run function in a new goroutine
		go checkLink(link, ch)
	}

	fmt.Println(<-ch)
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		ch <- "down"
		return
	}

	ch <- "up"
	fmt.Printf("%s is up\n", link)

}
