package runes

import "io"

var _ Writer = ByteWriter{}

// ByteWriter is a runes.Writer that writes to any given io.Writer.
type ByteWriter struct {
	io.Writer
}

// Write to the underlying io.Writer the string bytes for the given runes.
func (w ByteWriter) Write(p []rune) (n int, err error) {
	return w.Writer.Write([]byte(string(p)))
}
