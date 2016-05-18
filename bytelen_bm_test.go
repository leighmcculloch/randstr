package randstr

import "testing"

func BenchmarkByteLen0bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0x0)
	}
}

func BenchmarkByteLen8bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0xFF)
	}
}

func BenchmarkByteLen16bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0xFFFF)
	}
}

func BenchmarkByteLen32bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0xFFFFFFFF)
	}
}

func BenchmarkByteLen48bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0xFFFFFFFFFFFF)
	}
}

func BenchmarkByteLen64bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		byteLen(0x7FFFFFFFFFFFFFFF)
	}
}
