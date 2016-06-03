package randstr_test

import (
	"crypto/rand"
	"github.com/leighmcculloch/randstr/lib/randstr"
)

func ExampleString() {
	_ = randstr.String(rand.Reader, 50, randstr.ASCII)
}
