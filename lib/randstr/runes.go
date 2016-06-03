package randstr

type Charset interface {
	At(i int) rune
	Length() int
}

type CharsetArray []rune

func (c CharsetArray) At(i int) rune {
	return c[i]
}

func (c CharsetArray) Length() int {
	return len(c)
}

type CharsetRange struct {
	First rune
	Last  rune
}

func (c CharsetRange) At(i int) rune {
	return c.First + rune(i)
}

func (c CharsetRange) Length() int {
	return int(c.Last) - int(c.First) + 1
}

type Charsets []Charset

func (charsets Charsets) At(i int) rune {
	offset := 0
	for _, c := range charsets {
		length := offset + c.Length()
		if i < length {
			return c.At(i - offset)
		}
		offset = length
	}
	return 0
}

func (charsets Charsets) Length() int {
	length := 0
	for _, c := range charsets {
		length += c.Length()
	}
	return length
}
