# randstr
[![Linux/OSX Build Status](https://img.shields.io/travis/leighmcculloch/randstr.svg?label=linux%20%26%20osx)](https://travis-ci.org/leighmcculloch/randstr)
[![Windows Build Status](https://img.shields.io/appveyor/ci/leighmcculloch/randstr.svg?label=windows)](https://ci.appveyor.com/project/leighmcculloch/randstr)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/randstr.svg)](https://codecov.io/gh/leighmcculloch/randstr)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/randstr)](https://goreportcard.com/report/github.com/leighmcculloch/randstr)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/leighmcculloch/randstr)

Randstr is a package and a CLI that generates random strings (e.g. passwords), supporting unicode and emojis.

## Install

### Linux/Mac

```
curl -o /usr/local/bin/randstr https://raw.githubusercontent.com/leighmcculloch/randstr/binaries/$(uname -s | tr '[DL]' '[dl]')/amd64/randstr
chmod +x /usr/local/bin/randstr
```

or

```
brew install https://raw.githubusercontent.com/leighmcculloch/randstr/master/releases/homebrew/randstr.rb
```

### Windows

[Download the executable](https://raw.githubusercontent.com/leighmcculloch/randstr/binaries/windows/amd64/randstr.exe), and save it in your path.

### Source

```
go install github.com/leighmcculloch/randstr/cmd/randstr
```

## Usage

```
$ randstr
Ne*!Z|us'VRu;waO53_g{%*gwbY,vY\nw\wR/"5^cg1JkJ`k`l
```

On Mac, you can pipe the output directly to the clipboard with:

```
$ randstr | pbcopy
```

## Package

Use the package by go getting and importing:

```shell
go get github.com/leighmcculloch/randstr/lib/randstr
```

```go
import (
  "crypto/rand"
  "github.com/leighmcculloch/randstr/lib/randstr"
)

func main() {
  password := randstr.String(rand.Reader, randstr.ASCIICharset, 50)
  fmt.Printf("Password: %s\n", password)
}
```
