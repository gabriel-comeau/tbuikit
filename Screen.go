package termbox-uikit

import (
	"github.com/nsf/termbox-go"
)

// Basically a widget holder at the moment
type Screen struct {
	widgets           []Widget
	screenKeyBindings map[interface{}]EventCallback
	active            bool
}

// Sets this screen to active
func (this *Screen) Activate() {
	this.active = true
}

// Sets this screen to active
func (this *Screen) Deactivate() {
	this.active = false
}

// Predicate to see whether or not this screen is active.
func (this *Screen) IsActive() bool {
	return this.active
}

// Adds a new widget to the screen
func (this *Screen) AddWidget(widget Widget) {
	this.widgets = append(this.widgets, widget)
}

// Add a keybinding to the screen -- these override widget level keybindings
// so don't add keys here unless you're sure that the containing widgets won't
// need them.
//
// This function is for metakeys (those of type termbox.Key) - nonprinting keys basically
func (this *Screen) AddSpecialKeyCallback(event termbox.Key, callback EventCallback) {
	if len(this.screenKeyBindings) == 0 {
		this.screenKeyBindings = make(map[interface{}]EventCallback)
	}
	this.screenKeyBindings[event] = callback
}

// Add a keybinding to the screen -- these override widget level keybindings
// so don't add keys here unless you're sure that the containing widgets won't
// need them.
//
// This function is for printing keys - everything that normally prints a
// character to the screen.  It takes runes as it's argument.
func (this *Screen) AddCharKeyCallback(char rune, callback EventCallback) {
	if len(this.screenKeyBindings) == 0 {
		this.screenKeyBindings = make(map[interface{}]EventCallback)
	}
	this.screenKeyBindings[char] = callback
}

// To call when the screen resizes - it cycles through each widget
// and forces them to recalculate
func (this *Screen) DoResize() {
	for _, w := range this.widgets {
		w.CalculateSize()
	}
}

// Loop through our widgets and draw them all to the screen.
func (this *Screen) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, w := range this.widgets {
		w.Draw()
	}
	termbox.Flush()
}

// Iterate through the widgets and return the selected one
// or nil if there isn't one.  Also returns position this widget was
func (this *Screen) GetCurrentSelectedWidget() (w Widget, ind int) {
	for i, w := range this.GetSelectableWidgets() {
		if w.IsSelected() {
			return w, i
		}
	}
	return nil, 0
}

// Get all of the selectable widgets held by the screen
func (this *Screen) GetSelectableWidgets() []Widget {
	toReturn := make([]Widget, 0)
	for _, w := range this.widgets {
		if w.IsSelectable() {
			toReturn = append(toReturn, w)
		}
	}
	return toReturn
}

// Cycles through the selectable widgets
func (this *Screen) SelectNextWidget() {
	selectable := this.GetSelectableWidgets()
	prevSelected, prevPosition := this.GetCurrentSelectedWidget()
	count := len(selectable)
	if count > 1 {
		if (prevPosition + 1) > (count - 1) {
			// Start back over at 0
			prevSelected.Unselect()
			nextSelected := selectable[0]
			nextSelected.Select()
		} else {
			prevSelected.Unselect()
			nextSelected := selectable[prevPosition+1]
			nextSelected.Select()
		}
	}
}

// Keyboard event handling.  All calls to screen level
// callbacks get  a pointer to the screen in question.
func (this *Screen) HandleEvents(event interface{}) {

	currentWidget, _ := this.GetCurrentSelectedWidget()

	// Check for screen level keybindings
	if this.screenKeyBindings[event] != nil {
		this.screenKeyBindings[event](this, event)
	} else if currentWidget != nil {
		currentWidget.HandleEvents(event)
	}
}
