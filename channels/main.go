package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://stackoverflow.com/",
		"https://pythoninstitute.org/",
		"https://www.google.com/",
		"https://go.dev/",
	}

	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, "is up!")
}
