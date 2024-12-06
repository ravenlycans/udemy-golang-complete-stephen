package main

import (
	"fmt"
	"net/http"
	"time"
)

type chanMessage struct {
	message   string
	link      string
	nTimesRun int
	lastRan   time.Time
}

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

func checkLink(url string, ntr int, c chan chanMessage) {
	ntr++
	if fetch(url) {
		cm := chanMessage{message: fmt.Sprintf("Link %s is up, times checked: %d\n", url, ntr), link: url, nTimesRun: ntr, lastRan: time.Now()}
		c <- cm
	} else {
		cm := chanMessage{message: fmt.Sprintf("Link %s is down, times checked: %d\n", url, ntr), link: url, nTimesRun: ntr, lastRan: time.Now()}
		c <- cm
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

	c := make(chan chanMessage)

	fmt.Println("Starting to check links")
	fmt.Println("*******************************")

	for _, link := range links {
		go func() {
			checkLink(link, 0, c)
		}()
	}

	rDone := 0

	for l := range c {
		if len(l.message) > 0 {
			fmt.Printf("%s", l.message)
		}

		//		fmt.Printf("Last Ran: %s, Time Now: %s\n", l.lastRan.Format(time.RFC1123), time.Now().Format(time.RFC1123))

		if l.nTimesRun == 5 {
			rDone++
		} else {
			go func() {
				time.Sleep(time.Second * 5)
				checkLink(l.link, l.nTimesRun, c)
			}()
		}

		if rDone >= len(links) {
			break
		}
	}

	fmt.Println("*******************************")
	fmt.Println("Finished checking links.")
}
