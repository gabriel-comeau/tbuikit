package termbox-uikit

// This type wraps a string slice to be used
// to display strings on the screen.
// Has a maximum size and will truncate old strings once this is reached
type StringBuffer struct {
	holder   []string
	capacity int
}

// Call this to setup the slice when creating one of these
func (this *StringBuffer) Prepare(capacity int) {
	this.holder = make([]string, 0)
	this.capacity = capacity
}

// Adds a new element to the end of the stack, and will clip
// anything over capacity off the bottom.
func (this *StringBuffer) Add(str string) {
	this.holder = append(this.holder, str)
	if len(this.holder) > this.capacity {
		this.truncateOld()
	}
}

// Clear the buffer's contents.
func (this *StringBuffer) Clear() {
	this.holder = make([]string, 0)
}

// Gets the lines out of the buffer.  This will split any lines that are too long
// into two (or as many as it takes until they are shorter than line length) lines
// and then returns the last <lineCount> number of lines.
func (this *StringBuffer) GetContents(lineLength, lineCount int) []string {
	colStringCount := len(this.holder)
	if lineLength == 0 && colStringCount == 0 {
		return nil
	}
	cutOff := 0
	if lineCount < colStringCount {
		cutOff = colStringCount - lineCount
	}

	// Get the "messages" - which can be longer than a line
	unsplitLines := this.holder[cutOff:colStringCount]
	splitLines := make([]string, 0)
	// Add everything, splitting any lines that are too long into smaller lines
	// We'll rechop after
	for _, str := range unsplitLines {
		if len(str) <= lineLength {
			splitLines = append(splitLines, str)
		} else {
			newlySplitLines := SplitBufferLines(str, lineLength)
			for _, newLine := range newlySplitLines {
				splitLines = append(splitLines, newLine)
			}
		}
	}

	// now we need to make sure we haven't got more lines than lineCount
	if len(splitLines) > lineCount {
		splitLines = splitLines[(len(splitLines) - lineCount):]
	}

	return splitLines

}

// Clears oldest colorized strings to get back to capacity
func (this *StringBuffer) truncateOld() {
	length := len(this.holder)
	cutOff := length - this.capacity
	this.holder = this.holder[cutOff:length]
}
