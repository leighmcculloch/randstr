package charset

import "errors"

// Charsets is a charset that is the combination of other charsets.
type Charsets []Charset

// At returns the character at index i, by considering the internal charsets as a contiguous single charset.
func (cs Charsets) At(i int) rune {
	offset := 0
	for _, c := range cs {
		length := offset + c.Length()
		if i < length {
			return c.At(i - offset)
		}
		offset = length
	}
	panic(errors.New("runtime error: index out of range"))
}

// Length returns the number of characters in the charset, by summing the lengths of the internal charsets.
func (cs Charsets) Length() int {
	length := 0
	for _, c := range cs {
		length += c.Length()
	}
	return length
}
