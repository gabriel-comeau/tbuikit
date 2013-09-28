package tbuikit

import (
	"github.com/nsf/termbox-go"
)

// A widget which essentially knows how to draw message objects.
// It relies on a buffer type which meets the TextBuffer interface being passed in.
//
// To control where it is drawn on the screen, it uses a function value which returns the four corners
// of a rectangle, which represents it's location on the screen.
//
// These shouldn't be created via new() - use the CreateColorizedTextWidget() call instead.
type ColorizedStringWidget struct {
	rect         *Rectangle
	textColor    termbox.Attribute
	borderColor  termbox.Attribute
	bgColor      termbox.Attribute
	calcFunction CalcFunction
	buffer       *ColorizedStringBuffer
}

// This is the draw call - it takes a buffer type which meets the text buffer interface
// and draws the text in it to the screen at the positions defined by its rectangle.
func (this *ColorizedStringWidget) Draw() {
	if this.rect == nil {
		this.CalculateSize()
	}

	this.drawBorderAndBg()
	lines := this.buffer.GetContents(this.rect.Width()-1, this.rect.Height()-1)
	linesLen := len(lines)
	heightMod := 0
	for i := 0; i < linesLen; i++ {
		TermboxPrintf(this.rect.X1+1, this.rect.Y2-linesLen+heightMod, lines[i].Color, this.bgColor, lines[i].Text)
		heightMod++
	}
}

// This draws the border lines around the widget
//
// TODO: probably paint BG colors, thought this may need to be
// in tandem with the normal Draw function as well.
func (this *ColorizedStringWidget) drawBorderAndBg() {

	// Draw corners
	termbox.SetCell(this.rect.X1, this.rect.Y1, 0x250C, this.borderColor, this.bgColor)
	termbox.SetCell(this.rect.X2, this.rect.Y1, 0x2510, this.borderColor, this.bgColor)
	termbox.SetCell(this.rect.X1, this.rect.Y2, 0x2514, this.borderColor, this.bgColor)
	termbox.SetCell(this.rect.X2, this.rect.Y2, 0x2518, this.borderColor, this.bgColor)

	for i := this.rect.X1 + 1; i < this.rect.X2; i++ {
		termbox.SetCell(i, this.rect.Y1, 0x2500, this.borderColor, this.bgColor)
		termbox.SetCell(i, this.rect.Y2, 0x2500, this.borderColor, this.bgColor)
	}

	for i := this.rect.Y1 + 1; i < this.rect.Y2; i++ {
		termbox.SetCell(this.rect.X1, i, 0x2502, this.borderColor, this.bgColor)
		termbox.SetCell(this.rect.X2, i, 0x2502, this.borderColor, this.bgColor)
	}
}

// This widget cannot ever be selectable, so always return false.
func (this *ColorizedStringWidget) IsSelectable() bool {
	return false
}

// This widget cannot ever be selectable, so always return false.
func (this *ColorizedStringWidget) IsSelected() bool {
	return false
}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.
func (this *ColorizedStringWidget) Select() {}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.s
func (this *ColorizedStringWidget) Unselect() {}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.
func (this *ColorizedStringWidget) HandleEvents(event interface{}) {}

// Meant to be called when the terminal dimensions are resized, it calls the callback function and
// refigures out the sizing and positioning of the rectacngle.
func (this *ColorizedStringWidget) CalculateSize() {
	rect := CreateRectangle(this.calcFunction())
	this.rect = rect
}

// A "constructor" function to create new widgets.
func CreateColorizedTextWidget(textColor, borderColor, bgColor termbox.Attribute, calcFunction CalcFunction, buffer *ColorizedStringBuffer) *ColorizedStringWidget {
	widget := new(ColorizedStringWidget)
	widget.textColor = textColor
	widget.bgColor = bgColor
	widget.borderColor = borderColor
	widget.calcFunction = calcFunction
	widget.buffer = buffer
	return widget
}
