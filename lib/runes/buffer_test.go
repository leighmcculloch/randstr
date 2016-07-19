package runes

import "testing"

func TestNilBuffer(t *testing.T) {
	expectedRunes := []rune(nil)

	var b *Buffer

	runes := b.Runes()
	if len(runes) != len(expectedRunes) {
		t.Fatalf("Runes() = %#v, expected %#v", runes, expectedRunes)
	}
	for i, r := range runes {
		if r != runes[i] {
			t.Fatalf("Runes() = %#v, expected %#v", runes, expectedRunes)
		}
	}

	t.Logf("Runes() = %#v", runes)
}

func TestBuffer(t *testing.T) {
	expectedRunes := []rune("Hello ✈ Writer")

	b := Buffer{}

	var n int
	var err error

	n, err = b.Write([]rune("Hello"))
	if n == 5 && err == nil {
		t.Logf("Write(%#v) = %d, %#v", "Hello", n, err)
	} else {
		t.Fatalf("Write(%#v) = %d, %#v, expected %d, %#v", "Hello", n, err, 5, nil)
	}
	n, err = b.Write([]rune(" ✈ "))
	if n == 3 && err == nil {
		t.Logf("Write(%#v) = %d, %#v", "Hello", n, err)
	} else {
		t.Fatalf("Write(%#v) = %d, %#v, expected %d, %#v", "Hello", n, err, 3, nil)
	}
	n, err = b.Write([]rune("Buffer"))
	if n == 6 && err == nil {
		t.Logf("Write(%#v) = %d, %#v", "Hello", n, err)
	} else {
		t.Fatalf("Write(%#v) = %d, %#v, expected %d, %#v", "Hello", n, err, 6, nil)
	}

	runes := b.Runes()
	if len(runes) != len(expectedRunes) {
		t.Fatalf("Runes() = %#v, expected %#v", runes, expectedRunes)
	}
	for i, r := range runes {
		if r != runes[i] {
			t.Fatalf("Runes() = %#v, expected %#v", runes, expectedRunes)
		}
	}

	t.Logf("Runes() = %#v", runes)
}
