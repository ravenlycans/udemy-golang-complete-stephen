package main

import (
	"fmt"
	"net/http"
	"sync"
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

func checkLink(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	if fetch(url) {
		fmt.Printf("Link %s is up\n", url)
	} else {
		fmt.Printf("Link %s is down\n", url)
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

	fmt.Println("Starting to check links")
	fmt.Println("*******************************")

	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go checkLink(link, &wg)
	}

	wg.Wait()

	fmt.Println("*******************************")
	fmt.Println("Finished checking links.")
}
