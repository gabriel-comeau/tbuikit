package tbuikit

// Interfaces used in the library

// Interface defining widgets - a function to draw them to the screen
// and one to call when the screen is resized.
type Widget interface {
	Draw()
	CalculateSize()

	// Handle widget selection
	IsSelectable() bool
	IsSelected() bool
	Select()
	Unselect()

	// Handle keyboard events -- this is an interface{} because
	// it can either be a termbox.Key or rune at the moment
	HandleEvents(interface{})
}
