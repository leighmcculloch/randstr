package randstr_test

import (
	"crypto/rand"
	"github.com/leighmcculloch/randstr/lib/randstr"
)

func ExampleString() {
	_ = randstr.String(rand.Reader, randstr.ASCIICharset, 50)
}
