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
	ASCIIUppercaseRunes  = RuneArray(ASCIIUppercaseChars)
	ASCIILowercaseRunes  = RuneArray(ASCIILowercaseChars)
	ASCIINumericRunes    = RuneArray(ASCIINumericChars)
	ASCIIWhitespaceRunes = RuneArray(ASCIIWhitespaceChars)
	ASCIISymbolRunes     = RuneArray(ASCIISymbolChars)
	ASCIIRunes           = RuneArray(ASCIIChars)
)
