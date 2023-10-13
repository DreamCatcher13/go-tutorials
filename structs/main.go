package main

import "fmt"

// definition

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//var alex person  >>  go will assign Zero value to the fields, "", 0, 0, bool==False
	jim := person{
		firstName: "Jim",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "jim@tower.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim   & > access to mmr address of jim
	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName // * > access to value by mmr adress
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
