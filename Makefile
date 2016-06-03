test:
	go test -cover ./...

cover:
	go test -coverprofile=coverprofile-libcharset ./lib/charset
	go tool cover -html=coverprofile-libcharset
	go test -coverprofile=coverprofile-librandstr ./lib/randstr
	go tool cover -html=coverprofile-librandstr

bench:
	go test -bench=. -run='xxx' ./lib/randstr

build:
	go build ./lib/charset ./lib/randstr
	go build ./cmd/randstr

install:
	go install -ldflags "-X main.version=+`git rev-parse --short HEAD`" ./cmd/randstr

clean:
	go clean
	rm -fR binaries

release:
	mkdir -p binaries/linux/amd64
	mkdir -p binaries/darwin/amd64
	mkdir -p binaries/windows/amd64
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o binaries/linux/amd64/randstr ./cmd/randstr
	GOARCH=amd64 GOOS=darwin go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o binaries/darwin/amd64/randstr ./cmd/randstr
	GOARCH=amd64 GOOS=windows go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o binaries/windows/amd64/randstr.exe ./cmd/randstr
	
	git branch -D binaries 2>/dev/null | true
	git branch -D binaries-draft 2>/dev/null | true
	git checkout -b binaries-draft
	git add -f binaries
	git commit -m "Release to branch `binaries`"
	git subtree split --prefix binaries -b binaries
	git push --force origin binaries:binaries

	mkdir -p homebrew/Formula
	@echo "Homebrew formula for [randstr](https://github.com/leighmcculloch/randstr)." > homebrew/README.md
	@echo "\
	class Randstr < Formula \n\
		homepage \"https://github.com/leighmcculloch/randstr\" \n\
		url \"https://raw.githubusercontent.com/leighmcculloch/randstr/binaries/darwin/amd64/randstr\" \n\
		version \"$$(./binaries/darwin/amd64/randstr -shortversion)\" \n\
		sha256 \"$$(shasum -b -a 256 ./binaries/darwin/amd64/randstr |cut -f 1 -d " ")\" \n\
	\n\
		def install \n\
			bin.install \"randstr\" \n\
		end \n\
	\n\
		test do \n\
			system \"#{bin}/randstr\", \"-version\" \n\
		end \n\
	end" > homebrew/Formula/randstr.rb 
	git branch -D homebrew 2>/dev/null | true
	git branch -D homebrew-draft 2>/dev/null | true
	git checkout -b homebrew-draft
	git add -f homebrew
	git commit -m "Release to branch `master` of remote `homebrew`"
	git subtree split --prefix homebrew -b homebrew
	git push --force homebrew homebrew:master
	
	git checkout @{-2}
