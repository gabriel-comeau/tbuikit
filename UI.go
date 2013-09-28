package tbuikit

import (
	"github.com/nsf/termbox-go"
	"time"
)

// Represents the top level of a TermboxUIKit UI.  It manages
// screens, handles the global level keyboard (or resize) events,
// and handles shutting the entire UI down.
type UI struct {
	screenHolder      []*Screen
	globalKeyBindings map[interface{}]EventCallback
	uiShutdownChan    chan bool
}

func (this *UI) Start(quitChan chan bool) {
	if this.getActiveScreen() == nil {
		panic("No active screen")
	}

	// Make sure to initialize the channel
	this.uiShutdownChan = make(chan bool)

	this.mainLoop(quitChan)
}

// Sends the shutdown signal, breaking the main loop
func (this *UI) Shutdown() {
	this.uiShutdownChan <- true
}

// Adds a screen to the collection
func (this *UI) AddScreen(screen *Screen) {
	this.screenHolder = append(this.screenHolder, screen)
}

// Adds a global level event binding for a meta key
func (this *UI) AddSpecialKeyCallback(event termbox.Key, callback EventCallback) {
	if len(this.globalKeyBindings) == 0 {
		this.globalKeyBindings = make(map[interface{}]EventCallback)
	}
	this.globalKeyBindings[event] = callback
}

// Adds a global level event binding for a printable key
func (this *UI) AddCharKeyCallback(char rune, callback EventCallback) {
	if len(this.globalKeyBindings) == 0 {
		this.globalKeyBindings = make(map[interface{}]EventCallback)
	}
	this.globalKeyBindings[char] = callback
}

// Internal method for getting the active screen of the UI.
// Returns nil if nothing comes back as active.
func (this *UI) getActiveScreen() *Screen {
	var screen *Screen
	for _, s := range this.screenHolder {
		if s.IsActive() {
			screen = s
		}
	}
	return screen
}

// Performs the UI's main loop.  Fires up the termbox ui,
// creates an event queue, starts polling it for events and then
// hands off those events to the mapped handlers, starting at top level
// (this ui object) and moving down to the screen and widget level through
// handlers.
func (this *UI) mainLoop(quitChan chan bool) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	eventQueue := make(chan termbox.Event)

	// Read termbox events async on channel
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	// Take care of exiting from the entire UI here
	go func() {
		for {
			quitSig := <-this.uiShutdownChan
			if quitSig {

				// Shutdown the termbox UI
				termbox.Close()

				// Send the quit signal back to the ui caller
				quitChan <- true
			}
		}
	}()

	// Main event handling loop
	for {
		select {
		case ev := <-eventQueue:

			// Handle resize event
			if ev.Type == termbox.EventResize {
				this.getActiveScreen().DoResize()
			}

			// Check for top level keybindings
			// Calls the appropriate callback and passes an instance of the ui to it

			if ev.Type == termbox.EventKey && this.globalKeyBindings[ev.Key] != nil {
				this.globalKeyBindings[ev.Key](this, ev.Key)
			} else if ev.Type == termbox.EventKey && this.globalKeyBindings[ev.Ch] != nil {
				this.globalKeyBindings[ev.Ch](this, ev.Ch)
			} else if ev.Type == termbox.EventKey && ev.Key != 0 {
				this.getActiveScreen().HandleEvents(ev.Key)
			} else if ev.Type == termbox.EventKey && ev.Ch != 0 {
				this.getActiveScreen().HandleEvents(ev.Ch)
			}

		default:
			this.getActiveScreen().Draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
