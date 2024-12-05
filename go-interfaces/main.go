package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printSpanishGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Greetings!!"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating an spanish greeting
	return "Hola!!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printSpanishGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}
