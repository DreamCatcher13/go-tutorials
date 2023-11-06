package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface { // => new custom type bot
	getGreeting() string // any type in this program and which have a func getGreating() string
} // that type is also bot

type logWriter struct {
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	resp, err := http.Get("https://google.com/")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// READER intf
	/*	bs := make([]byte, 99999) // Read func fill the byte slice until it is full, so declare it big
		resp.Body.Read(bs)
		fmt.Println(string(bs)) */

	// WRITER intf
	// io.Copy(os.Stdout, resp.Body)  // dst with Writer(),  src with Reader()
	lw := logWriter{}
	io.Copy(lw, resp.Body)
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

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs)) // print to stdOut
	fmt.Println("Number of bytes:", len(bs))
	return len(bs), nil
}
