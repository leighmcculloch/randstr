package randstr

import "math/bits"

// byteLen returns the minimum number of bytes that would be required to store the integer.
func byteLen(n int) int {
	return (bits.Len(uint(n))-1)/8 + 1
}
