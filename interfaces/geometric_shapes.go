package main

import (
	"interfaces/shape"
	"log"
)

type FlatShaper interface {
	Area() float64
	Perimeter() float64
}

func main() {
	myRectangle := shape.Rectangle{
		Base:   10.0,
		Height: 3.0,
	}
	calculate(&myRectangle)

	myCircle := &shape.Circle{Radio: 5.0}
	calculate(myCircle)
}

// calculate Prints in logs the perimeter and area of a plain geometrical shape
func calculate(fs FlatShaper) {
	log.Printf("The shape Perimeter is: %f", fs.Perimeter())
	log.Printf("The shape Area is: %.2f", fs.Area())
}
