package randstr

import "testing"

func TestByteLen(t *testing.T) {
	tests := []struct {
		input  int
		output int
	}{
		{0, 1},
		{1, 1},
		{2, 1},
		{0xFF, 1},
		{0xFF + 1, 2},
		{0xFFFF, 2},
		{0xFFFF + 1, 3},
		{0xFFFFFF, 3},
		{0xFFFFFF + 1, 4},
		{0xFFFFFFFF, 4},
		{0xFFFFFFFF + 1, 5},
		{0xFFFFFFFFFF, 5},
		{0xFFFFFFFFFF + 1, 6},
		{0xFFFFFFFFFFFF, 6},
		{0xFFFFFFFFFFFF + 1, 7},
		{0xFFFFFFFFFFFFFF, 7},
		{0xFFFFFFFFFFFFFF + 1, 8},
		{0x7FFFFFFFFFFFFFFF, 8},
		{-0, 1},
		{-1, 8},
		{-2, 8},
		{-0xFF, 8},
		{-0xFF + 1, 8},
		{-0xFFFF, 8},
		{-0xFFFF + 1, 8},
		{-0xFFFFFF, 8},
		{-0xFFFFFF + 1, 8},
		{-0xFFFFFFFF, 8},
		{-0xFFFFFFFF + 1, 8},
		{-0xFFFFFFFFFF, 8},
		{-0xFFFFFFFFFF + 1, 8},
		{-0xFFFFFFFFFFFF, 8},
		{-0xFFFFFFFFFFFF + 1, 8},
		{-0xFFFFFFFFFFFFFF, 8},
		{-0xFFFFFFFFFFFFFF + 1, 8},
		{-0x7FFFFFFFFFFFFFFF, 8},
	}

	for _, test := range tests {
		bytelen := byteLen(test.input)
		if bytelen == test.output {
			t.Logf("ByteLen(0x%016X) is %d", test.input, bytelen)
		} else {
			t.Errorf("ByteLen(0x%016X) is %d, expected %d", test.input, bytelen, test.output)
		}
	}
}
