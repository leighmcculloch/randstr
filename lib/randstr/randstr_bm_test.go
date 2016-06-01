package randstr

import "testing"

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(50, ASCIIRunes)
	}
}
