package termbox-uikit

import (
	"github.com/nsf/termbox-go"
)

// Defines simple, method-less types

// Callback function used to dynamically resize the rectangle
// when the screen changes size
type CalcFunction func() (x1, x2, y1, y2 int)

// Represents positions - typically where
// text is inside a widget.
type ScreenPosition int

// Holds a string and a color to print it in
type ColorizedString struct {
	Color termbox.Attribute
	Text  string
}

// A callback function which can be mapped to a key in various
// levels of the interface.  When the keybinding is detected,
// the callback is executed and passed a pointer to the ui element
// which was called and the key which was pressed.
//
// The key pressed can either be a rune (for printable keys) or
// a meta character of type termbox.Key
type EventCallback func(interface{}, interface{})
