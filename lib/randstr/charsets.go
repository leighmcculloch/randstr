package randstr

var (
	ASCIIUppercase  = CharsetArray("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	ASCIILowercase  = CharsetArray("abcdefghijklmnopqrstuvwxyz")
	ASCIINumeric    = CharsetArray("0123456789")
	ASCIIWhitespace = CharsetArray(" ")
	ASCIISymbol     = CharsetArray("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
	ASCII           = CharsetArray(string(ASCIIUppercase) + string(ASCIILowercase) + string(ASCIIWhitespace) + string(ASCIINumeric) + string(ASCIISymbol))
)
