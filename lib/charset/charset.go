package charset

// Charset is an interface that provides a set of characters.
type Charset interface {
	// At returns the character at the zero based index i.
	At(i int) rune
	// Length returns the number of characters within the charset.
	Length() int
}
