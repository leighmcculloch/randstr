package randstr

import "testing"

func TestRunesAt(t *testing.T) {
	tests := []struct {
		Runes        Runes
		Index        int
		ExpectedRune rune
	}{
		{RuneArray{'a', 'b', 'c'}, 0, 'a'},
		{RuneArray{'a', 'b', 'c'}, 1, 'b'},
		{RuneArray{'a', 'b', 'c'}, 2, 'c'},
		{RuneRange{'a', 'c'}, 0, 'a'},
		{RuneRange{'a', 'c'}, 1, 'b'},
		{RuneRange{'a', 'c'}, 2, 'c'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}}), 0, 'a'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}}), 1, 'b'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}}), 2, 'c'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 0, 'a'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 1, 'b'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 2, 'c'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 3, 'x'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 4, 'y'},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 5, 'z'},
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
		{RuneRange{}, 1},
		{RuneRange{'a', 'a'}, 1},
		{RuneRange{'a', 'c'}, 3},
		{MultipleRunes{}, 0},
		{MultipleRunes([]Runes{RuneRange{'a', 'a'}}), 1},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}}), 3},
		{MultipleRunes([]Runes{RuneRange{'a', 'c'}, RuneRange{'x', 'z'}}), 6},
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
