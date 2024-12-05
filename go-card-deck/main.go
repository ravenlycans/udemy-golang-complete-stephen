package main

import (
	"fmt"
	"math/rand"
	"time"
)

func section03Challenge() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	ooe := r.Intn(100)

	fmt.Printf("Number to test is: %d\n", ooe)

	if ooe%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

func main() {
	section03Challenge()
}
