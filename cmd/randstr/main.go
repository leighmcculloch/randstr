package main

import (
	"flag"
	"fmt"

	"github.com/leighmcculloch/randstr"
)

var version string

var printHelp bool
var printVersion bool
var length int
var chars string

func init() {
	flag.BoolVar(&printHelp, "help", false, "display this usage")
	flag.BoolVar(&printVersion, "version", false, "display the version")
	flag.IntVar(&length, "l", 50, "`length` of the string generated")
	flag.StringVar(&chars, "c", randstr.ASCIIChars, "`characters` to potentially use in the string, supporting unicode and emojis")
}

func main() {
	flag.Parse()

	if printVersion {
		fmt.Printf("randstr version %s\n", version)
		return
	}

	if printHelp {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		return
	}

	fmt.Println(randstr.New(length, []rune(chars)))
}
