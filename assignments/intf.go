package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	tr := triangle{
		height: 20,
		base:   10,
	}

	sq := square{
		sideLength: 5,
	}

	printArea(tr)
	printArea(sq)

}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println("The area of the shape is:", s.getArea())
}
