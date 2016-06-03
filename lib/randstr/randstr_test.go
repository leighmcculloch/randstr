package randstr

import (
	"crypto/rand"
	"github.com/leighmcculloch/randstr/lib/charset"
	"io"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		Rand    io.Reader
		Charset charset.Charset
		Length  int
	}{
		{rand.Reader, charset.ASCII, 50},
		{rand.Reader, charset.UnicodeEmoji, 50},
		{rand.Reader, charset.CharsetRange{0, 0x7FFF}, 50},
	}
	attempts := 500
	for _, test := range tests {
		seen := make(map[string]bool)
		for i := 0; i < attempts; i++ {
			s := String(test.Rand, test.Charset, test.Length)
			sLen := len([]rune(s))
			switch {
			case sLen != test.Length:
				t.Fatalf("Created string %s len %d, expected %d.", s, sLen, test.Length)
			case seen[s]:
				t.Fatalf("Created string %s len %d, which is not unique.", s, sLen)
			default:
				t.Logf("Created string %s len %d, which is unique so far.", s, sLen)
			}
			seen[s] = true
		}
		t.Logf("Created %d strings, all unique, all len %d.", len(seen), test.Length)
	}
}
