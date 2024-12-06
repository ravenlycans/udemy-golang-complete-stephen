package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	r, err := http.Get("http://www.google.com")

	if err != nil {
		panic(err)
	}

	bs := make([]byte, 1024)
	bs, err = io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Printf("Read %d bytes\n", len(bs))
}
