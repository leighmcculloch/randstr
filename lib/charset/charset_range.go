package charset

import "errors"

// CharsetRange is a contiguous range of characters starting at First and ending at and including Last.
type CharsetRange struct {
	First rune
	Last  rune
}

// At returns the character at index i.
func (c CharsetRange) At(i int) rune {
	r := c.First + rune(i)
	if r > c.Last {
		panic(errors.New("runtime error: index out of range"))
	}
	return r
}

// Length returns the number of characters in the charset.
func (c CharsetRange) Length() int {
	return int(c.Last) - int(c.First) + 1
}
