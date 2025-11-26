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

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("%s is down\n", link)
		return
	}
	fmt.Printf("%s is up\n", link)
}
