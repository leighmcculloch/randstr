package randstr_test

import (
	"crypto/rand"

	"4d63.com/randstr/lib/charset"
	"4d63.com/randstr/lib/randstr"
)

func ExampleString() {
	_, _ = randstr.String(rand.Reader, charset.ASCII, 50)
}
