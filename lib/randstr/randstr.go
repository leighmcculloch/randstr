package randstr

import (
	"io"
)

// Rune returns a random array of runes of length given, using only the runes given.
func Rune(rand io.Reader, length int, charset Charset) []rune {
	r := make([]rune, length)

	runesCount := charset.Length()
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		io.ReadFull(rand, bytes)

		n := bytesToInt(bytes)
		if n < 0 {
			n = -n
		}
		n = n % runesCount

		r[i] = charset.At(n)
	}

	return r
}

// String returns a random string of length given, using only the runes given.
func String(rand io.Reader, length int, charset Charset) string {
	return string(Rune(rand, length, charset))
}
