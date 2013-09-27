package termbox-uikit

// Represents a simple rectangle
type Rectangle struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

// Gets the rectangle's width
func (this *Rectangle) Width() int {
	return this.X2 - this.X1
}

// Gets the rectangle's height
func (this *Rectangle) Height() int {
	return this.Y2 - this.Y1
}

// Creates a new rectangle object
func CreateRectangle(x1, x2, y1, y2 int) *Rectangle {
	rect := new(Rectangle)
	rect.X1 = x1
	rect.X2 = x2
	rect.Y1 = y1
	rect.Y2 = y2
	return rect
}
