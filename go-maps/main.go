package main

import "fmt"

func main() {
	colors := make(map[string]string)
	colors["red"] = "#FF0000"
	colors["green"] = "#00FF00"
	colors["blue"] = "#0000FF"
	colors["white"] = "#FFFFFF"
	colors["black"] = "#000000"
	colors["weird"] = "#4BF745"

	printMap(colors)
}

func printMap(m map[string]string) {
	for col, hex := range m {
		fmt.Printf("Hex code for %s is %s\n", col, hex)
	}
}
