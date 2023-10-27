package main

import "fmt"

type bot interface { // => new custom type bot
	getGreeting() string // any type in this program and which have a func getGreating() string
} // that type is also bot

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic
	return "Howdy, stranger!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic
	// we are not using var sb, so we can remove it
	return "HOLA!"
}
