package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"strings"

	"github.com/leighmcculloch/randstr/lib/charset"
	"github.com/leighmcculloch/randstr/lib/randstr"
)

var version string

var charsetOptions = map[string]charset.Charset{
	"ASCIIUppercase":  charset.ASCIIUppercase,
	"ASCIILowercase":  charset.ASCIILowercase,
	"ASCIINumeric":    charset.ASCIINumeric,
	"ASCIISpace":      charset.ASCIISpace,
	"ASCII":           charset.ASCII,
	"UnicodePassword": charset.UnicodePassword,
	"UnicodeEmoji":    charset.UnicodeEmoji,
}

var charsetOptionNames = func() []string {
	names := []string{}
	for name := range charsetOptions {
		names = append(names, name)
	}
	return names
}()

var printHelp bool
var printVersion bool
var printShortVersion bool
var length int
var chars string
var charsetNameList string

func init() {
	flag.BoolVar(&printHelp, "help", false, "display this usage")
	flag.BoolVar(&printVersion, "version", false, "display the version")
	flag.BoolVar(&printShortVersion, "shortversion", false, "print the version to stdout")
	flag.IntVar(&length, "l", 50, "`length` of the string generated")
	flag.StringVar(&chars, "chars", "", "`chars` to use in the string, supporting unicode and emojis")
	flag.StringVar(&charsetNameList, "charset", "ASCII", fmt.Sprintf("comma separated list of `charsets` to use in the string, e.g. %s", strings.Join(charsetOptionNames, ",")))
}

func main() {
	flag.Parse()

	if printShortVersion {
		fmt.Println(version)
		return
	}

	if printVersion {
		fmt.Printf("randstr version %s\n", version)
		return
	}

	if printHelp {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	var ch charset.Charset
	if chars != "" {
		ch = charset.CharsetArray(chars)
	} else {
		charsetNames := strings.Split(charsetNameList, ",")
		charsets := charset.Charsets{}
		for _, name := range charsetNames {
			charsetOption, ok := charsetOptions[name]
			if !ok {
				fmt.Printf("Error: Charset %s unknown. See the help.\n", name)
				continue
			}
			charsets = append(charsets, charsetOption)
		}
		ch = charsets
	}

	if ch.Length() == 0 {
		return
	}

	randomString, err := randstr.String(rand.Reader, ch, length)
	if err != nil {
		fmt.Printf("Error encountered:\n%s", err.Error())
		return
	}
	fmt.Println(randomString)
}
