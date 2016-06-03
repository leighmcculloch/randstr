package randstr

import (
	"crypto/rand"
	"testing"
)

func BenchmarkNewSmallSource(b *testing.B) {
	runes := RuneArray(ASCIIRunes)
	for i := 0; i < b.N; i++ {
		New(rand.Reader, 50, runes)
	}
}

func BenchmarkNewLargeSource(b *testing.B) {
	runes := RuneRange{0, 9999999}
	for i := 0; i < b.N; i++ {
		New(rand.Reader, 50, runes)
	}
}

func BenchmarkNewMixedSource(b *testing.B) {
	runes := MultipleRunes([]Runes{
		RuneArray(ASCIISymbolRunes),
		RuneArray(ASCIINumericRunes),
		RuneArray(ASCIIWhitespaceRunes),
		RuneRange{'a', 'z'},
		RuneRange{'A', 'Z'},
	})
	for i := 0; i < b.N; i++ {
		New(rand.Reader, 50, runes)
	}
}
