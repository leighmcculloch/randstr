package randstr

import "testing"

func BenchmarkNewSmallSource(b *testing.B) {
	runes := RuneArray(ASCIIRunes)
	for i := 0; i < b.N; i++ {
		New(50, runes)
	}
}

func BenchmarkNewLargeSource(b *testing.B) {
	runesLit := [9999999]rune{}
	runes := RuneArray(runesLit[:])
	for i := 0; i < b.N; i++ {
		New(50, runes)
	}
}
