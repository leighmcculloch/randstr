package randstr

const (
	ASCIIUppercaseChars  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	ASCIILowercaseChars  = "abcdefghijklmnopqrstuvwxyz"
	ASCIINumericChars    = "0123456789"
	ASCIIWhitespaceChars = " "
	ASCIISymbolChars     = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	ASCIIChars           = ASCIIUppercaseChars + ASCIILowercaseChars + ASCIIWhitespaceChars + ASCIINumericChars + ASCIISymbolChars
)

var (
	ASCIIUppercaseRunes  = []rune(ASCIIUppercaseChars)
	ASCIILowercaseRunes  = []rune(ASCIILowercaseChars)
	ASCIINumericRunes    = []rune(ASCIINumericChars)
	ASCIIWhitespaceRunes = []rune(ASCIIWhitespaceChars)
	ASCIISymbolRunes     = []rune(ASCIISymbolChars)
	ASCIIRunes           = []rune(ASCIIChars)
)
