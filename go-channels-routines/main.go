package main

import (
	"fmt"
	"net/http"
)

func checkLink(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://golang.org",
	}

	fmt.Println("Starting to check links")
	fmt.Println("*******************************")

	for _, link := range links {
		if checkLink(link) {
			fmt.Printf("Link %s is up\n", link)
		} else {
			fmt.Printf("Link %s is down\n", link)
		}
	}
	fmt.Println("*******************************")
	fmt.Println("Finished checking links.")
}
