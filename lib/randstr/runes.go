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
