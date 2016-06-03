package randstr

// bytesToInt returns the integer stored in the bytes array interpreted in big endian order.
func bytesToInt(bytes []byte) int {
	n := 0
	for _, b := range bytes {
		n = (n << 8) | int(b)
	}
	return n
}
