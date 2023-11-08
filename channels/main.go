package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://stackoverflow.com/",
		"https://pythoninstitute.org/",
		"https://www.google.com/",
		"https://go.dev/",
	}

	c := make(chan string) // creating a channel for communication between go routines
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c { // wait for a value in channel and assign it to "l"
		go func(link string) { // func literal => lambda => anonymous func
			time.Sleep(5 * time.Second)
			checkLink(link, c) // receive a value from a channel;  blocking code
		}(l) // call that literal
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link // sending data to channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
