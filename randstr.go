package randstr

import (
	"crypto/rand"
	"io"
)

const (
	ASCIIUppercaseChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ASCIILowercaseChars  = "abcdefghijklmnopqrstuvwxyz"
	ASCIINumericChars    = "0123456789"
	ASCIIWhitespaceChars = " "
	ASCIISymbolChars     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	ASCIIChars           = ASCIIUppercaseChars + ASCIILowercaseChars + ASCIIWhitespaceChars + ASCIINumericChars + ASCIISymbolChars
)

var (
	ASCIIUppercaseRunes  = []rune(ASCIIUppercaseChars)
	ASCIILowercaseRunes  = []rune(ASCIILowercaseChars)
	ASCIINumericRunes    = []rune(ASCIINumericChars)
	ASCIIWhitespaceRunes = []rune(ASCIIWhitespaceChars)
	ASCIISymbolRunes     = []rune(ASCIISymbolChars)
	ASCIIRunes           = []rune(ASCIIChars)
)

// New generates a new random string of length given, using only the characters given in runes.
func New(length int, runes []rune) string {
	password := make([]rune, length)

	runesCount := len(runes)
	byteCount := byteLen(runesCount)
	bytes := make([]byte, byteCount)

	for i := 0; i < length; i++ {
		io.ReadFull(rand.Reader, bytes)

		n := bytesToInt(bytes)
		if n < 0 {
			n = -n
		}
		n = n % runesCount

		password[i] = runes[n]
	}

	return string(password)
}
