package randstr

import (
	"crypto/rand"
	"testing"
)

func BenchmarkNewSmallSource(b *testing.B) {
	charsets := CharsetArray(ASCIICharset)
	for i := 0; i < b.N; i++ {
		String(rand.Reader, 50, charsets)
	}
}

func BenchmarkNewLargeSource(b *testing.B) {
	charsets := CharsetRange{0, 9999999}
	for i := 0; i < b.N; i++ {
		String(rand.Reader, 50, charsets)
	}
}

func BenchmarkNewMixedSource(b *testing.B) {
	charsets := Charsets([]Charset{
		CharsetArray(ASCIISymbolCharset),
		CharsetArray(ASCIINumericCharset),
		CharsetArray(ASCIIWhitespaceCharset),
		CharsetRange{'a', 'z'},
		CharsetRange{'A', 'Z'},
	})
	for i := 0; i < b.N; i++ {
		String(rand.Reader, 50, charsets)
	}
}
