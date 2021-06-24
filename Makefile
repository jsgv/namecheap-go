all: test buildcommand

build:
	go build -o dist/namecheap cmd/namecheap/main.go

test:
	go test -v ./...

