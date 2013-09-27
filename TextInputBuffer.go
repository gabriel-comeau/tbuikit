package termbox-uikit

// This buffer represents the storage for any field a user can type text into
type TextInputBuffer struct {
	charHolder []rune
	length     int
}

// Adds a new element to the end of the stack (just a method form of append)
func (this *TextInputBuffer) Add(char rune) {
	// 0 is unlimited length
	if this.length == 0 {
		this.charHolder = append(this.charHolder, char)
	} else if this.length > 1 && len(this.charHolder) < this.length {
		this.charHolder = append(this.charHolder, char)
	}
}

// Removes the last element from the buffer
func (this *TextInputBuffer) Backspace() {
	length := len(this.charHolder)
	if length > 1 {
		this.charHolder = this.charHolder[0 : length-1]
	} else {
		this.charHolder = make([]rune, 0)
	}
}

// Wraps the call to toString and then clear,
// which is what the enter key should do
func (this *TextInputBuffer) ReturnAndClear() string {
	contents := string(this.charHolder)
	this.Clear()
	return contents
}

// Clears the buffer
func (this *TextInputBuffer) Clear() {
	this.charHolder = make([]rune, 0)
}

func (this *TextInputBuffer) SetLength(length int) {
	this.length = length
}

// Returns the text contents of this buffer as a slice of strings.  The number
// of strings returned depends on how long the text the buffer contains
// and how long the desired lineLength is.  lineLength can be set to 0
// in order to force the buffer's contents into a single string and return
// a single element slice containing it.
//
// Line count can optionally be used to return only the last n lines,
// if a maximum number of lines is a concern.
func (this *TextInputBuffer) GetLines(lineLength, lineCount int) []string {
	var lines []string
	stringified := string(this.charHolder)

	if lineLength != 0 && len(stringified) > lineLength {
		lines = SplitBufferLines(stringified, lineLength)
		if lineCount != 0 {
			totalLen := len(lines)
			lines = lines[totalLen-lineCount-1:]
		}
	} else {
		lines = make([]string, 1)
		lines[0] = stringified
	}
	return lines
}

// Checks if this buffer is empty
func (this *TextInputBuffer) IsEmpty() bool {
	if len(this.charHolder) == 0 {
		return true
	}
	return false
}
