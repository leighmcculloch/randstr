package randstr

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"

	"github.com/leighmcculloch/randstr/lib/charset"
)

func TestString(t *testing.T) {
	tests := []struct {
		Rand    io.Reader
		Charset charset.Charset
		Length  int
	}{
		{rand.Reader, charset.ASCII, 50},
		{rand.Reader, charset.UnicodeEmoji, 50},
		{rand.Reader, charset.CharsetRange{First: 0, Last: 0x7FFF}, 50},
	}
	attempts := 500
	for _, test := range tests {
		seen := make(map[string]bool)
		for i := 0; i < attempts; i++ {
			s, err := String(test.Rand, test.Charset, test.Length)
			if err != nil {
				t.Fatalf("Creating string returned error: %#v", err)
			}
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

type ReaderThatErrors struct{}

func (r ReaderThatErrors) Read(_ []byte) (int, error) {
	return 0, fmt.Errorf("Error occurred during read")
}

func TestString_Error(t *testing.T) {
	expectedErr := fmt.Errorf("Error occurred during read")
	rand := ReaderThatErrors{}
	_, err := String(rand, charset.ASCII, 50)
	if err == nil {
		t.Fatalf("Creating string returned without error, expected error: %#v", expectedErr)
	}
	if err.Error() != expectedErr.Error() {
		t.Fatalf("Creating string returned error %#v, expected error: %#v", err, expectedErr)
	}
	t.Logf("Creating string returned error: %#v", err)
}
