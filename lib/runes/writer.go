package runes

// Writer is an interface that writes runes.
type Writer interface {
	// Write writes runes and returns the number of runes written or error if an error occurred.
	Write(p []rune) (n int, err error)
}
