package main

type Circle struct {
	Radius float64
}
type Square struct {
	Length float64
}
type Rectangle struct {
	Length float64
	Width  float64
}

type Shape interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (s Square) Area() float64 {
	return s.Length * s.Length
}
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func testInterface() {
	c := Circle{Radius: 5}
	s := Square{Length: 5}
	r := Rectangle{Length: 5, Width: 10}

	shapes := []Shape{c, s, r}

	for _, shape := range shapes {
		println(shape.Area())
	}
}
