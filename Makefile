buildcommand:
	go build -o dist/namecheap cmd/namecheap/main.go

test:
	go test ./pkg/cmd/
