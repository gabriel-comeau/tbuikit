package tbuikit

import (
	"github.com/nsf/termbox-go"
)

// A widget which prints text to the screen, like a button but deliberately can't
// be selected.
type LabelWidget struct {
	labelText    string
	textPosition ScreenPosition
	drawBorders  bool

	textColor   termbox.Attribute
	borderColor termbox.Attribute
	bgColor     termbox.Attribute

	calcFunction CalcFunction
	rect         *Rectangle
}

// Draw the label every iteration of the main loop.  Figure out where to put the button text within the label,
// make sure the borders get drawn and then draw the text in the figured out location.
func (this *LabelWidget) Draw() {
	if this.rect == nil {
		this.CalculateSize()
	}

	if this.drawBorders {
		this.drawBorderAndBg()
	}

	// Decide where the label text should be drawn.  Keep in mind
	// the Termboxprintf function still prints a string left-to-right,
	// so the X coord is the first rune of the string and the it advances
	// to the right.
	var x, y int

	if this.textPosition == TOP_LEFT {
		y = this.rect.Y1 + 1
		x = this.rect.X1 + 1
	} else if this.textPosition == TOP_RIGHT {
		y = this.rect.Y1 + 1
		x = this.rect.X2 - len(this.labelText) // - 1?
	} else if this.textPosition == BOTTOM_LEFT {
		y = this.rect.Y2 - 1
		x = this.rect.X1 + 1
	} else if this.textPosition == BOTTOM_RIGHT {
		y = this.rect.Y2 - 1
		x = this.rect.X2 - len(this.labelText)
	} else {
		// default to center
		y = this.rect.Y2 - (this.rect.Height() / 2)
		x = this.rect.X2 - (this.rect.Width() / 2) - (len(this.labelText) / 2)
	}

	TermboxPrintf(x, y, this.textColor, this.bgColor, this.labelText)
}

// This draws the border lines around the widget
//
// TODO: probably paint BG colors, thought this may need to be
// in tandem with the normal Draw function as well.
func (this *LabelWidget) drawBorderAndBg() {

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

// Meant to be called when the terminal dimensions are resized, it calls the callback function and
// refigures out the sizing and positioning of the rectacngle.
func (this *LabelWidget) CalculateSize() {
	rect := CreateRectangle(this.calcFunction())
	this.rect = rect
}

// This widget cannot ever be selectable, so always return false.
func (this *LabelWidget) IsSelectable() bool {
	return false
}

// This widget cannot ever be selectable, so always return false.
func (this *LabelWidget) IsSelected() bool {
	return false
}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.
func (this *LabelWidget) Select() {}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.s
func (this *LabelWidget) Unselect() {}

// This kind of widget cannot be selected.
// For the moment do nothing, this is just here to satisfy the interface.
func (this *LabelWidget) HandleEvents(event interface{}) {}

// Setter for the label's displayed text.
func (this *LabelWidget) SetText(text string) {
	this.labelText = text
}

// A "constructor" function to create new widgets.
func CreateLabelWidget(text string, drawBor bool, textPos ScreenPosition, textCol, bgCol, borCol termbox.Attribute,
	calcFunction CalcFunction) *LabelWidget {

	widget := new(LabelWidget)
	widget.labelText = text
	widget.textPosition = textPos
	widget.drawBorders = drawBor

	widget.textColor = textCol
	widget.bgColor = bgCol
	widget.borderColor = borCol

	widget.calcFunction = calcFunction

	return widget
}
