package charset

// CharsetArray is an array of characters.
type CharsetArray []rune

// At returns the character at index i.
func (c CharsetArray) At(i int) rune {
	return c[i]
}

// Length returns the number of characters in the charset.
func (c CharsetArray) Length() int {
	return len(c)
}
