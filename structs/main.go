package main

import "fmt"

// definition
type person struct {
	firstName string
	lastName  string
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person // go will assign Zero value to the fields, "", 0, 0, bool==False
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)
}
