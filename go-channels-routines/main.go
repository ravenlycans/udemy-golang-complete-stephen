package main

import (
	"fmt"
	"net/http"
)

func fetch(url string) bool {
	resp, err := http.Get(url)

	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func checkLink(url string, c chan string) {
	if fetch(url) {
		c <- fmt.Sprintf("Link %s is up\n", url)
	} else {
		c <- fmt.Sprintf("Link %s is down\n", url)
	}
}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://go.dev",
		"https://golang.org",
	}

	c := make(chan string)

	fmt.Println("Starting to check links")
	fmt.Println("*******************************")

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Printf("%s", <-c)
	}

	fmt.Println("*******************************")
	fmt.Println("Finished checking links.")
}
