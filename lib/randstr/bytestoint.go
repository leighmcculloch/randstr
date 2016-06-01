package randstr

func bytesToInt(bytes []byte) int {
	n := 0
	for _, b := range bytes {
		n = (n << 8) | int(b)
	}
	return n
}
