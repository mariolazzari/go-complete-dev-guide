package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error getting google.com", err)
		os.Exit(1)
	}
	fmt.Println(resp)
}
