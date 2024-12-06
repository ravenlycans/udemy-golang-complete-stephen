package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type logWriter struct{}

func main() {
	r, err := http.Get("http://www.google.com")

	if err != nil {
		panic(err)
	}

	/*
		bs := make([]byte, 1024)
		bs, err = io.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}
	*/

	nw, ew := io.Copy(logWriter{}, r.Body)
	if ew != nil {
		panic(ew)
	}
	fmt.Printf("Read %d bytes\n", nw)

	//	fmt.Println(string(bs))
	//	fmt.Printf("Read %d bytes\n", len(bs))
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Printf("LogWriter (%s): %s\n", time.Now().Format(time.RFC822), string(bs))

	return len(bs), nil
}
