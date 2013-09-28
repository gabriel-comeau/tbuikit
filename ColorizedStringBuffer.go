package tbuikit

// This type wraps a string slice to be used
// to display string-wrapper objects on the screen.
// The objects are simple tuples, containing the string and the color
// to display it as.
// Has a maximum size and will truncate old strings once this is reached
type ColorizedStringBuffer struct {
	holder   []*ColorizedString
	capacity int
}

// Call this to setup the slice when creating one of these
func (this *ColorizedStringBuffer) Prepare(capacity int) {
	this.holder = make([]*ColorizedString, 0)
	this.capacity = capacity
}

// Adds a new element to the end of the stack, and will clip
// anything over capacity off the bottom.
func (this *ColorizedStringBuffer) Add(msg *ColorizedString) {
	this.holder = append(this.holder, msg)
	if len(this.holder) > this.capacity {
		this.truncateOld()
	}
}

// Clear the buffer's contents.
func (this *ColorizedStringBuffer) Clear() {
	this.holder = make([]*ColorizedString, 0)
}

// Gets the lines out of the buffer.  This will split any lines that are too long
// into two (or as many as it takes until they are shorter than line length) lines
// and then returns the last <lineCount> number of lines.
func (this *ColorizedStringBuffer) GetContents(lineLength, lineCount int) []*ColorizedString {
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
	splitLines := make([]*ColorizedString, 0)
	// Add everything, splitting any lines that are too long into smaller lines
	// We'll rechop after
	for _, colString := range unsplitLines {
		if len(colString.Text) <= lineLength {
			splitLines = append(splitLines, colString)
		} else {
			newlySplitLines := SplitBufferLines(colString.Text, lineLength)
			for _, newLine := range newlySplitLines {
				cs := new(ColorizedString)
				cs.Text = newLine
				cs.Color = colString.Color
				splitLines = append(splitLines, cs)
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
func (this *ColorizedStringBuffer) truncateOld() {
	length := len(this.holder)
	cutOff := length - this.capacity
	this.holder = this.holder[cutOff:length]
}
