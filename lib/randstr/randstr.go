package randstr

import (
	"crypto/rand"
	"io"
)

// New generates a new random string of length given, using only the characters given in runes.
func New(length int, runes Runes) string {
	s := make([]rune, length)

	runesCount := runes.Length()
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		io.ReadFull(rand.Reader, bytes)

		n := bytesToInt(bytes)
		if n < 0 {
			n = -n
		}
		n = n % runesCount

		s[i] = runes.At(n)
	}

	return string(s)
}
