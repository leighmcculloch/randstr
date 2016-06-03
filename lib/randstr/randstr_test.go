package randstr

import (
	"crypto/rand"
	"testing"
)

func TestString(t *testing.T) {
	length := 50
	attempts := 500
	passwordsSeen := make(map[string]bool)
	for i := 0; i < attempts; i++ {
		password := String(rand.Reader, length, RuneArray(ASCIINumericRunes))
		passwordLen := len(password)
		switch {
		case passwordLen != length:
			t.Fatalf("Created password %s len %d, expected %d.", password, passwordLen, length)
		case passwordsSeen[password]:
			t.Fatalf("Created password %s len %d, which is not unique.", password, passwordLen)
		default:
			t.Logf("Created password %s len %d, which is unique so far.", password, passwordLen)
		}
		passwordsSeen[password] = true
	}
	t.Logf("Created %d passwords, all unique, all len %d.", len(passwordsSeen), length)
}
