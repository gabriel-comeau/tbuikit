package tbuikit

import (
	"github.com/nsf/termbox-go"
)

// A widget meant for password input.  It stores the real value in the buffer but it displays
// everything as an "*"
//
// To control where it is drawn on the screen, it uses a function value which returns the four corners
// of a rectangle, which represents it's location on the screen.
//
// These shouldn't be created via new() - use the CreatePasswordInputWidget() call instead.
type PasswordInputWidget struct {
	rect              *Rectangle
	hasCursor         bool
	defaultTextColor  termbox.Attribute
	defaultBgColor    termbox.Attribute
	selectedBgColor   termbox.Attribute
	calcFunction      CalcFunction
	buffer            *TextInputBuffer
	isSelectable      bool
	selected          bool
	widgetKeyBindings map[interface{}]EventCallback
	defaultHandler    bool
}

// This is the draw call - it takes a buffer type which meets the text buffer interface
// and draws the text in it to the screen at the positions defined by its rectangle.
func (this *PasswordInputWidget) Draw() {
	if this.rect == nil {
		this.CalculateSize()
	}

	this.drawBorderAndBg()

	lines := this.buffer.GetLines(this.rect.Width()-1, this.rect.Height()-1)
	linesLen := len(lines)
	heightMod := 0
	for i := 0; i < linesLen; i++ {
		astLen := len(lines[i])
		asts := ""
		for j := 0; j < astLen; j++ {
			asts += "*"
		}
		TermboxPrintf(this.rect.X1+1, this.rect.Y2-linesLen+heightMod, this.defaultTextColor, termbox.ColorDefault, asts)
		heightMod++
	}

	if this.hasCursor && this.selected {
		if this.buffer.IsEmpty() {
			termbox.SetCursor(this.rect.X1+1, this.rect.Y1+(this.rect.Height())-1)
		} else {
			MoveCursor(this.rect.X1, this.rect.Y2-1, lines[linesLen-1])
		}
	}
}

// This draws the border lines around the widget
//
// TODO: probably paint BG colors, thought this may need to be
// in tandem with the normal Draw function as well.
func (this *PasswordInputWidget) drawBorderAndBg() {

	var color termbox.Attribute
	if this.selected {
		color = this.selectedBgColor | termbox.AttrBold
	} else {
		color = this.defaultBgColor
	}

	// Draw corners
	termbox.SetCell(this.rect.X1, this.rect.Y1, 0x250C, color, termbox.ColorDefault)
	termbox.SetCell(this.rect.X2, this.rect.Y1, 0x2510, color, termbox.ColorDefault)
	termbox.SetCell(this.rect.X1, this.rect.Y2, 0x2514, color, termbox.ColorDefault)
	termbox.SetCell(this.rect.X2, this.rect.Y2, 0x2518, color, termbox.ColorDefault)

	for i := this.rect.X1 + 1; i < this.rect.X2; i++ {
		termbox.SetCell(i, this.rect.Y1, 0x2500, color, termbox.ColorDefault)
		termbox.SetCell(i, this.rect.Y2, 0x2500, color, termbox.ColorDefault)
	}

	for i := this.rect.Y1 + 1; i < this.rect.Y2; i++ {
		termbox.SetCell(this.rect.X1, i, 0x2502, color, termbox.ColorDefault)
		termbox.SetCell(this.rect.X2, i, 0x2502, color, termbox.ColorDefault)
	}
}

// Meant to be called when the terminal dimensions are resized, it calls the callback function and
// refigures out the sizing and positioning of the rectacngle.
func (this *PasswordInputWidget) CalculateSize() {
	rect := CreateRectangle(this.calcFunction())
	this.rect = rect
}

// Check if this widget should be flaggable as selected.
func (this *PasswordInputWidget) IsSelectable() bool {
	return this.isSelectable
}

// Check if this widget is flagged as selected.  Accessor
// because eventually want to implement logic to test for isSelectable
func (this *PasswordInputWidget) IsSelected() bool {
	return this.selected
}

// Selects this widget - it'd probably make sense
// to return an error if this widget isn't selectable
func (this *PasswordInputWidget) Select() {
	if this.isSelectable {
		this.selected = true
	}
}

// Unset selection status
func (this *PasswordInputWidget) Unselect() {
	this.selected = false
}

// Take widget level printable-key (rune) handler function
func (this *PasswordInputWidget) AddCharKeyCallback(char rune, callback EventCallback) {
	this.widgetKeyBindings[char] = callback
}

// Take widget level meta-key (termbox.Key) handler function
func (this *PasswordInputWidget) AddSpecialKeyCallback(key termbox.Key, callback EventCallback) {
	this.widgetKeyBindings[key] = callback
}

// Enable using the default key bindings for the widget.
func (this *PasswordInputWidget) UseDefaultKeys(use bool) {
	this.defaultHandler = use
}

// Get the buffer - if we're able to switch selected items
// we'll probably want to be able to switch the buffer at the
// same time.
func (this *PasswordInputWidget) GetBuffer() *TextInputBuffer {
	return this.buffer
}

// If this widget is selected, handle key inputs based on mapped keys
func (this *PasswordInputWidget) HandleEvents(event interface{}) {
	if this.widgetKeyBindings[event] != nil {
		this.widgetKeyBindings[event](this, event)
	} else {
		if this.defaultHandler {
			this.handleDefaultKeys(event)
		}
	}
}

// This method handles the typical keys passed into a text input widget.
// Printable characters, backspace/delete and spacebar.
//
// Enter is deliberately left out - different applications will likely
// have radically different notions about what should be done with the buffers
// content and when of course.
//
// Additionally, this handler function needs to be explicitly enabled - not
// enabling it means that the application will have to manually define key/char
// bindings for it's event handler.
func (this *PasswordInputWidget) handleDefaultKeys(event interface{}) {

	key, ok := event.(termbox.Key)
	if ok {
		if key == termbox.KeySpace {
			this.GetBuffer().Add(' ')
		} else if key == termbox.KeyBackspace || key == termbox.KeyBackspace2 {
			this.GetBuffer().Backspace()
		}
	} else {
		char, charOk := event.(rune)
		if charOk {
			if char != ' ' {
				this.GetBuffer().Add(char)
			}
		}
	}
}

// A "constructor" function to create new widgets.
func CreatePasswordInputWidget(hasCursor bool, color termbox.Attribute, bg termbox.Attribute, selbg termbox.Attribute,
	calcFunction CalcFunction, buffer *TextInputBuffer, selectable bool, selected bool) *PasswordInputWidget {

	widget := new(PasswordInputWidget)
	widget.hasCursor = hasCursor
	widget.defaultTextColor = color
	widget.defaultBgColor = bg
	widget.selectedBgColor = selbg
	widget.calcFunction = calcFunction
	widget.buffer = buffer
	widget.selected = selected
	widget.isSelectable = selectable

	widget.widgetKeyBindings = make(map[interface{}]EventCallback)

	return widget
}
