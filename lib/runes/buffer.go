package runes

var _ Writer = &Buffer{}

// Buffer buffers runes.
type Buffer struct {
	buffer []rune
}

// Write the runes to the buffer. Never errors.
func (b *Buffer) Write(p []rune) (n int, err error) {
	b.buffer = append(b.buffer, p...)
	return len(p), nil
}

// Runes returns the
func (b *Buffer) Runes() []rune {
	if b == nil {
		return nil
	}
	return b.buffer[:]
}
