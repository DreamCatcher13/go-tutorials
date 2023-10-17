package main

import "fmt"

func main() {
	//var colors map[string]string
	//colors := make(map[string]string)
	//delete(colors, "white")  // map and a key

	colors := map[string]string{ // [key]value of type string
		"red":   "#ff0000",
		"green": "#4df754",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code of color", color, "is", hex)
	}
}
