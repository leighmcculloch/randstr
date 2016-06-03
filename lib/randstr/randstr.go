package randstr

import (
	"github.com/leighmcculloch/randstr/lib/charset"
	"io"
)

// Runes returns a random array of runes of length, using only the runes given and the io.Reader as the source of randomness.
func runes(rand io.Reader, charset charset.Charset, length int) []rune {
	r := make([]rune, length)

	runesCount := charset.Length()
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		io.ReadFull(rand, bytes)

		n := bytesToInt(bytes)
		n = n % runesCount

		r[i] = charset.At(n)
	}

	return r
}

// String returns a random string of length, using only the runes given and the io.Reader as the source of randomness. For cryptographic randomness use crypto/rand's Reader.
func String(rand io.Reader, charset charset.Charset, length int) string {
	return string(runes(rand, charset, length))
}
