package strtmethintr

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Calculate area of rectange", func(t *testing.T) {
		rectangle := Rectangle{15.0, 10.0}
		want := 150.0
		checkArea(t, rectangle, want)
	})
	t.Run("Calculate area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func TestAreaTable(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	areaTests := []struct {
		shape Shape
		want  float64
		name  string
	}{
		{Rectangle{20.0, 10.0}, 200, "Calculate area of rectange"},
		{Circle{1.0}, 3.141592653589793, "Calculate area of circle"},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}

}
