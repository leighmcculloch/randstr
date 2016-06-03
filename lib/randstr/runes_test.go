package randstr

import "testing"

func TestCharsetsAt(t *testing.T) {
	tests := []struct {
		Charset      Charset
		Index        int
		ExpectedRune rune
	}{
		{CharsetArray{'a', 'b', 'c'}, 0, 'a'},
		{CharsetArray{'a', 'b', 'c'}, 1, 'b'},
		{CharsetArray{'a', 'b', 'c'}, 2, 'c'},
		{CharsetRange{'a', 'c'}, 0, 'a'},
		{CharsetRange{'a', 'c'}, 1, 'b'},
		{CharsetRange{'a', 'c'}, 2, 'c'},
		{Charsets{CharsetRange{'a', 'c'}}, 0, 'a'},
		{Charsets{CharsetRange{'a', 'c'}}, 1, 'b'},
		{Charsets{CharsetRange{'a', 'c'}}, 2, 'c'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 0, 'a'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 1, 'b'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 2, 'c'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 3, 'x'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 4, 'y'},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 5, 'z'},
	}

	for _, test := range tests {
		r := test.Charset.At(test.Index)
		if r == test.ExpectedRune {
			t.Logf("%#v.At(%d) = %#v", test.Charset, test.Index, r)
		} else {
			t.Errorf("%#v.At(%d) = %#v, expected %#v", test.Charset, test.Index, r, test.ExpectedRune)
		}
	}
}

func TestCharsetsLength(t *testing.T) {
	tests := []struct {
		Charset        Charset
		ExpectedLength int
	}{
		{CharsetArray{}, 0},
		{CharsetArray{'a', 'b', 'c'}, 3},
		{CharsetRange{}, 1},
		{CharsetRange{'a', 'a'}, 1},
		{CharsetRange{'a', 'c'}, 3},
		{Charsets{}, 0},
		{Charsets{CharsetRange{'a', 'a'}}, 1},
		{Charsets{CharsetRange{'a', 'c'}}, 3},
		{Charsets{CharsetRange{'a', 'c'}, CharsetRange{'x', 'z'}}, 6},
	}

	for _, test := range tests {
		l := test.Charset.Length()
		if l == test.ExpectedLength {
			t.Logf("%#v.Length() = %d", test.Charset, l)
		} else {
			t.Errorf("%#v.Length() = %d, expected %d", test.Charset, l, test.ExpectedLength)
		}
	}
}
