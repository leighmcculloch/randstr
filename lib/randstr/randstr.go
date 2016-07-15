package randstr

import (
	"io"

	"github.com/leighmcculloch/randstr/lib/charset"
)

// Runes returns a random array of runes of length, using only the runes given and the io.Reader as the source of randomness.
func runes(rand io.Reader, charset charset.Charset, length int) ([]rune, error) {
	r := make([]rune, length)

	runesCount := charset.Length()
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		_, err := io.ReadFull(rand, bytes)
		if err != nil {
			return nil, err
		}

		n := bytesToInt(bytes)
		n = n % runesCount

		r[i] = charset.At(n)
	}

	return r, nil
}

// String returns a random string of length, using only the runes given and the io.Reader as the source of randomness. For cryptographic randomness use crypto/rand's Reader.
func String(rand io.Reader, charset charset.Charset, length int) (string, error) {
	r, err := runes(rand, charset, length)
	return string(r), err
}
