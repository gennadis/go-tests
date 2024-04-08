package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Trianlgle struct {
	Base   float64
	LegOne float64
	LegTwo float64
}

func (t Trianlgle) Area() float64 {
	halfPerimeter := t.Perimeter() / 2
	return math.Sqrt(halfPerimeter * (halfPerimeter - t.Base) * (halfPerimeter - t.LegOne) * (halfPerimeter - t.LegTwo))
}

func (t Trianlgle) Perimeter() float64 {
	return t.Base + t.LegOne + t.LegTwo
}
