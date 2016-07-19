package randstr

import (
	"io"

	"github.com/leighmcculloch/randstr/lib/charset"
	"github.com/leighmcculloch/randstr/lib/runes"
)

// Write writes to the writer random runes (as bytes), using only the runes given and the io.Reader as the source of randomness. For cryptographic randomness use crypto/rand's Reader.
func Write(w runes.Writer, rand io.Reader, charset charset.Charset, length int) error {
	runesCount := charset.Length()
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		if _, err := io.ReadFull(rand, bytes); err != nil {
			return err
		}

		n := bytesToInt(bytes)
		n = n % runesCount

		r := charset.At(n)
		if _, err := w.Write([]rune{r}); err != nil {
			return err
		}
	}

	return nil
}

// Runes returns a random array of runes, using only the runes given and the io.Reader as the source of randomness. For cryptographic randomness use crypto/rand's Reader.
func Runes(rand io.Reader, charset charset.Charset, length int) ([]rune, error) {
	b := runes.Buffer{}
	if err := Write(&b, rand, charset, length); err != nil {
		return nil, err
	}
	r := b.Runes()
	return r, nil
}

// String returns a random string of length, using only the runes given and the io.Reader as the source of randomness. For cryptographic randomness use crypto/rand's Reader.
func String(rand io.Reader, charset charset.Charset, length int) (string, error) {
	r, err := Runes(rand, charset, length)
	return string(r), err
}
