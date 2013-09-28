package tbuikit

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

// Basic functions

// Prints a string to a termbox buffer.
// Takes the height and starting x position and then prints the string RTL
func TermboxPrint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

// Same as print_tb but takes a formatted string
func TermboxPrintf(x, y int, fg, bg termbox.Attribute, format string, args ...interface{}) {
	s := fmt.Sprintf(format, args...)
	TermboxPrint(x, y, fg, bg, s)
}

// Gets only the termbox's width - since it constantly returns both values
// and we don't always want both, this is basically a wrapper around
// _ = h
func GetTermboxWidth() int {
	w, h := termbox.Size()
	_ = h
	return w
}

// Same as getTermboxWith but for height
func GetTermboxHeight() int {
	w, h := termbox.Size()
	_ = w
	return h
}

// Move the cursor to the end of the provided string,
// starting at a given xOffset (not all widgets start at
// the left edge of the screen!)
func MoveCursor(xOffset, y int, bufferLine string) {
	length := len(bufferLine)
	termbox.SetCursor(xOffset+length+1, y)
}

// Splits up a string into a slice of strings, making "lines" of
// text to display.  The criteria is to split by width (buffer widget width)
//
// TODO It'd be nice to split on whitespace instead of in the middle of a word!
func SplitBufferLines(stringToSplit string, width int) []string {
	lines := make([]string, 0)
	temp := ""
	j := 1
	for i := 0; i < len(stringToSplit); i++ {
		temp += string(stringToSplit[i])
		j++
		if j == width {
			j = 0
			lines = append(lines, temp)
			temp = ""
		} else if i == len(stringToSplit)-1 {
			lines = append(lines, temp)
		}
	}
	return lines
}
