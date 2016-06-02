test:
	go test -cover ./lib/randstr ./cmd/randstr

cover:
	go test -coverprofile=coverprofile ./lib/randstr
	go tool cover -html=coverprofile

bench:
	go test -bench=. -run='xxx' ./lib/randstr

build:
	go build ./lib/randstr
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
	cd cmd/randstr && GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o ../../binaries/linux/amd64/randstr .
	cd cmd/randstr && GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o ../../binaries/darwin/amd64/randstr .
	cd cmd/randstr && GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=+`git rev-parse --short HEAD`" -o ../../binaries/windows/amd64/randstr.exe .
	git branch -D binaries 2>/dev/null | true
	git branch -D binaries-draft 2>/dev/null | true
	git checkout -b binaries-draft && \
	git add -f binaries && \
	git commit -m "Release to branch `binaries`" && \
	git subtree split --prefix binaries -b binaries && \
	git push --force origin binaries:binaries && \
	git checkout -


