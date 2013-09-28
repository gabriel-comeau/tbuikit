package tbuikit

import (
	"github.com/nsf/termbox-go"
)

// A widget which acts like a button.  It can be selected and then "pressed"
// via a key binding.
type ButtonWidget struct {
	buttonText   string
	textPosition ScreenPosition

	defaultTextColor    termbox.Attribute
	selectedTextColor   termbox.Attribute
	defaultBorderColor  termbox.Attribute
	selectedBorderColor termbox.Attribute
	defaultBgColor      termbox.Attribute
	selectedBgColor     termbox.Attribute

	calcFunction      CalcFunction
	rect              *Rectangle
	isSelectable      bool
	selected          bool
	widgetKeyBindings map[interface{}]EventCallback
}

// Draw the button every iteration of the main loop.  Figure out where to put the button text within the button,
// make sure the borders get drawn and then draw the text in the figured out location.
func (this *ButtonWidget) Draw() {
	if this.rect == nil {
		this.CalculateSize()
	}

	this.drawBorderAndBg()

	// Decide where the button text should be drawn.  Keep in mind
	// the Termboxprintf function still prints a string left-to-right,
	// so the X coord is the first rune of the string and the it advances
	// to the right.
	var x, y int

	if this.textPosition == TOP_LEFT {
		y = this.rect.Y1 + 1
		x = this.rect.X1 + 1
	} else if this.textPosition == TOP_RIGHT {
		y = this.rect.Y1 + 1
		x = this.rect.X2 - len(this.buttonText) // - 1?
	} else if this.textPosition == BOTTOM_LEFT {
		y = this.rect.Y2 - 1
		x = this.rect.X1 + 1
	} else if this.textPosition == BOTTOM_RIGHT {
		y = this.rect.Y2 - 1
		x = this.rect.X2 - len(this.buttonText)
	} else {
		// default to center
		y = this.rect.Y2 - (this.rect.Height() / 2)
		x = this.rect.X2 - (this.rect.Width() / 2) - (len(this.buttonText) / 2)
	}

	if this.selected {
		TermboxPrintf(x, y, this.selectedTextColor, this.selectedBgColor, this.buttonText)
		termbox.HideCursor()
	} else {
		TermboxPrintf(x, y, this.defaultTextColor, this.defaultBgColor, this.buttonText)
	}
}

// This draws the border lines around the widget
//
// TODO: probably paint BG colors, thought this may need to be
// in tandem with the normal Draw function as well.
func (this *ButtonWidget) drawBorderAndBg() {

	var borderColor termbox.Attribute
	if this.selected {
		borderColor = this.selectedBorderColor | termbox.AttrBold
	} else {
		borderColor = this.defaultBorderColor
	}

	var bgColor termbox.Attribute
	if this.selected {
		bgColor = this.selectedBgColor
	} else {
		bgColor = this.defaultBgColor
	}

	// Draw corners
	termbox.SetCell(this.rect.X1, this.rect.Y1, 0x250C, borderColor, bgColor)
	termbox.SetCell(this.rect.X2, this.rect.Y1, 0x2510, borderColor, bgColor)
	termbox.SetCell(this.rect.X1, this.rect.Y2, 0x2514, borderColor, bgColor)
	termbox.SetCell(this.rect.X2, this.rect.Y2, 0x2518, borderColor, bgColor)

	for i := this.rect.X1 + 1; i < this.rect.X2; i++ {
		termbox.SetCell(i, this.rect.Y1, 0x2500, borderColor, bgColor)
		termbox.SetCell(i, this.rect.Y2, 0x2500, borderColor, bgColor)
	}

	for i := this.rect.Y1 + 1; i < this.rect.Y2; i++ {
		termbox.SetCell(this.rect.X1, i, 0x2502, borderColor, bgColor)
		termbox.SetCell(this.rect.X2, i, 0x2502, borderColor, bgColor)
	}
}

// Meant to be called when the terminal dimensions are resized, it calls the callback function and
// refigures out the sizing and positioning of the rectacngle.
func (this *ButtonWidget) CalculateSize() {
	rect := CreateRectangle(this.calcFunction())
	this.rect = rect
}

// Check if this widget should be flaggable as selected.
func (this *ButtonWidget) IsSelectable() bool {
	return this.isSelectable
}

// Check if this widget is flagged as selected.  Accessor
// because eventually want to implement logic to test for isSelectable
func (this *ButtonWidget) IsSelected() bool {
	return this.selected
}

// Selects this widget - it'd probably make sense
// to return an error if this widget isn't selectable
func (this *ButtonWidget) Select() {
	if this.isSelectable {
		this.selected = true
	}
}

// Unset selection status
func (this *ButtonWidget) Unselect() {
	this.selected = false
}

// Take widget level printable-key (rune) handler function
func (this *ButtonWidget) AddCharKeyCallback(char rune, callback EventCallback) {
	this.widgetKeyBindings[char] = callback
}

// Take widget level meta-key (termbox.Key) handler function
func (this *ButtonWidget) AddSpecialKeyCallback(key termbox.Key, callback EventCallback) {
	this.widgetKeyBindings[key] = callback
}

// If this widget is selected, handle key inputs based on mapped keys
func (this *ButtonWidget) HandleEvents(event interface{}) {
	if this.widgetKeyBindings[event] != nil {
		this.widgetKeyBindings[event](this, event)
	}
}

// Setter for the button's displayed text.
func (this *ButtonWidget) SetText(btnText string) {
	//TODO: Maybe do something about too-long text by
	// looking at the width of the rect?
	this.buttonText = btnText
}

// A "constructor" function to create new widgets.
func CreateButtonWidget(buttonText string, textPos ScreenPosition, defTextCol, selTextCol, defBgCol, selBgCol, defBorCol, selBorCol termbox.Attribute,
	calcFunction CalcFunction, selectable bool, selected bool) *ButtonWidget {

	widget := new(ButtonWidget)
	widget.buttonText = buttonText
	widget.textPosition = textPos

	widget.defaultTextColor = defTextCol
	widget.selectedTextColor = selTextCol
	widget.defaultBgColor = defBgCol
	widget.selectedBgColor = selBgCol
	widget.defaultBorderColor = defBorCol
	widget.selectedBorderColor = selBorCol

	widget.calcFunction = calcFunction
	widget.selected = selected
	widget.isSelectable = selectable

	widget.widgetKeyBindings = make(map[interface{}]EventCallback)

	return widget
}
