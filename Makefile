deps:
	go mod download
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest

generate: deps
	~/go/bin/swagger generate client --target=.

clean:
	rm -rf client

integration:
	go test ./... -tags=integration
