package randstr

// byteLen returns the minimum number of bytes that would be required to store the integer.
func byteLen(n int) int {
	b := 0
	s := 1
	for s != 0 {
		b++
		s = s << 8
		if n < s {
			return b
		}
	}
	return b
}
