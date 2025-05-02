package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("expected %.2f, got %.2f", expected, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()

		if got != expected {
			t.Errorf("%#v expected area is %.2f, got %.2f", shape, expected, got)
		}
	}

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12.0, 6.0}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 35.0},
	}
	
	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			checkArea(t, testCase.shape, testCase.hasArea)
		})
	}
}