package randstr

import (
	"crypto/rand"
	"testing"

	"4d63.com/randstr/lib/charset"
)

func BenchmarkNewSmallSource(b *testing.B) {
	charsets := charset.CharsetArray(charset.ASCII)
	for i := 0; i < b.N; i++ {
		String(rand.Reader, charsets, 50)
	}
}

func BenchmarkNewLargeSource(b *testing.B) {
	charsets := charset.CharsetRange{0, 9999999}
	for i := 0; i < b.N; i++ {
		String(rand.Reader, charsets, 50)
	}
}

func BenchmarkNewMixedSource(b *testing.B) {
	charsets := charset.Charsets([]charset.Charset{
		charset.CharsetArray(charset.ASCIISymbol),
		charset.CharsetArray(charset.ASCIINumeric),
		charset.CharsetArray(charset.ASCIISpace),
		charset.CharsetRange{'a', 'z'},
		charset.CharsetRange{'A', 'Z'},
	})
	for i := 0; i < b.N; i++ {
		String(rand.Reader, charsets, 50)
	}
}
