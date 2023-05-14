package shape

type Rectangle struct {
	Base, Height float64
}

func (r *Rectangle) Area() float64 {
	// A = b * a
	return r.Base * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	// P = 2(b + a)
	return 2 * (r.Base + r.Height)
}
