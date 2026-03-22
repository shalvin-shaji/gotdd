package strtmethintr

func Perimeter(rectange Rectange) (perimeter float64) {
	return 2 * (rectange.Width + rectange.Height)
}

func Area(rectange Rectange) (area float64) {
	return (rectange.Width * rectange.Height)
}
