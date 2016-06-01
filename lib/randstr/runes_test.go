package randstr

import "testing"

func TestRunesAt(t *testing.T) {
	tests := []struct {
		Runes        Runes
		Index        int
		ExpectedRune rune
	}{
		{RuneArray{'a', 'b', 'c'}, 1, 'b'},
	}

	for _, test := range tests {
		r := test.Runes.At(test.Index)
		if r == test.ExpectedRune {
			t.Logf("%#v.At(%d) = %#v", test.Runes, test.Index, r)
		} else {
			t.Errorf("%#v.At(%d) = %#v, expected %#v", test.Runes, test.Index, r, test.ExpectedRune)
		}
	}
}

func TestRunesLength(t *testing.T) {
	tests := []struct {
		Runes          Runes
		ExpectedLength int
	}{
		{RuneArray{}, 0},
		{RuneArray{'a', 'b', 'c'}, 3},
	}

	for _, test := range tests {
		l := test.Runes.Length()
		if l == test.ExpectedLength {
			t.Logf("%#v.Length() = %d", test.Runes, l)
		} else {
			t.Errorf("%#v.Length() = %d, expected %d", test.Runes, l, test.ExpectedLength)
		}
	}
}
