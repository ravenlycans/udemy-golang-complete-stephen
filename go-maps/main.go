package main

import "fmt"

func main() {
	//	var colors map[string]string

	//	colors := make(map[string]string)

	/*
		colors := map[string]string{
			"red":   "#FF0000",
			"green": "#00FF00",
			"blue":  "#0000FF",
			"weird": "#4BF745",
		}
	*/
	/*
		colors["red"] = "#FF0000"
		colors["green"] = "#00FF00"
		colors["blue"] = "#0000FF"
		colors["white"] = "#FFFFFF"
		colors["weird"] = "#4BF745"
	*/
	colors := make(map[int]string)

	colors[10] = "white"

	fmt.Println(colors)

	delete(colors, 10)

	fmt.Println(colors)
}
