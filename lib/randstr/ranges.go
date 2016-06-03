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
	ASCIIUppercase  = CharsetArray(ASCIIUppercaseChars)
	ASCIILowercase  = CharsetArray(ASCIILowercaseChars)
	ASCIINumeric    = CharsetArray(ASCIINumericChars)
	ASCIIWhitespace = CharsetArray(ASCIIWhitespaceChars)
	ASCIISymbol     = CharsetArray(ASCIISymbolChars)
	ASCII           = CharsetArray(ASCIIChars)
)
