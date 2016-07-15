package randstr_test

import (
	"crypto/rand"

	"github.com/leighmcculloch/randstr/lib/charset"
	"github.com/leighmcculloch/randstr/lib/randstr"
)

func ExampleString() {
	_, _ = randstr.String(rand.Reader, charset.ASCII, 50)
}
