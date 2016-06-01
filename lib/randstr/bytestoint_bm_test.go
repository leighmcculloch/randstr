package randstr

import "testing"

func BenchmarkBytesToInt0bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0x00})
	}
}

func BenchmarkBytesToInt8bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0xFF})
	}
}

func BenchmarkBytesToInt16bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0xFF, 0xFF})
	}
}

func BenchmarkBytesToInt32bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0xFF, 0xFF, 0xFF, 0xFF})
	}
}

func BenchmarkBytesToInt48bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	}
}

func BenchmarkBytesToInt64bit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytesToInt([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})
	}
}
