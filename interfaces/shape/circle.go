package shape

import "math"

type Circle struct {
	Radio float64
}

func (c *Circle) Area() float64 {
	// A = π * r²
	return math.Pi * math.Pow(c.Radio, 2)
}

func (c *Circle) Perimeter() float64 {
	// P = 2 * π * r
	return 2 * math.Pi * c.Radio
}
