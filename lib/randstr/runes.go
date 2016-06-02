package randstr

type Runes interface {
	At(i int) rune
	Length() int
}

type RuneArray []rune

func (r RuneArray) At(i int) rune {
	return r[i]
}

func (r RuneArray) Length() int {
	return len(r)
}

type RuneRange struct {
	First rune
	Last  rune
}

func (r RuneRange) At(i int) rune {
	return r.First + rune(i)
}

func (r RuneRange) Length() int {
	return int(r.Last) - int(r.First) + 1
}

type MultipleRunes []Runes

func (r MultipleRunes) At(i int) rune {
	offset := 0
	for _, runes := range r {
		length := offset + runes.Length()
		if i < length {
			return runes.At(i - offset)
		}
		offset = length
	}
	return 0
}

func (r MultipleRunes) Length() int {
	length := 0
	for _, runes := range r {
		length += runes.Length()
	}
	return length
}
