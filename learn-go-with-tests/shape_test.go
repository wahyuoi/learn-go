package learn_go_with_tests

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 7}, want: 84},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Width:3, Height:4}, want: 6.0},
	}

	for _, tt := range areaTests {
		t.Run(fmt.Sprintf("%v", tt), func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("got %g want %g", got, tt.want)
			}
		})
	}
}
