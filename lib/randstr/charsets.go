package randstr

var (
	ASCIIUppercaseCharset  = CharsetArray("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	ASCIILowercaseCharset  = CharsetArray("abcdefghijklmnopqrstuvwxyz")
	ASCIINumericCharset    = CharsetArray("0123456789")
	ASCIIWhitespaceCharset = CharsetArray(" ")
	ASCIISymbolCharset     = CharsetArray("!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
	ASCIICharset           = CharsetArray(ASCIIUppercaseCharset.String() + ASCIILowercaseCharset.String() + ASCIIWhitespaceCharset.String() + ASCIINumericCharset.String() + ASCIISymbolCharset.String())
)
