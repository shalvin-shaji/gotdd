package strtmethintr

func Perimeter(rectange Rectangle) (perimeter float64) {
	return 2 * (rectange.Width + rectange.Height)
}

func Area(rectange Rectangle) (area float64) {
	return (rectange.Width * rectange.Height)
}
