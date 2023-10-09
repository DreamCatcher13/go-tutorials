package main

import "fmt"

func main() {
	integers := []int{}

	for i := 0; i <= 10; i++ {
		integers = append(integers, i)
	}

	for _, n := range integers {
		if n%2 == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}

}
