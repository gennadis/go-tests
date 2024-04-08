package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Trianlgle{Base: 15, LegOne: 8, LegTwo: 17}, hasArea: 60.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}

}
func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape        Shape
		hasPerimeter float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36.0},
		{shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{shape: Trianlgle{Base: 15, LegOne: 8, LegTwo: 17}, hasPerimeter: 40.0},
	}
	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.hasPerimeter {
			t.Errorf("got %g want %g", got, tt.hasPerimeter)
		}
	}

}
