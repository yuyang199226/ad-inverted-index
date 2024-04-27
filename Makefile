build:
	go build -o bin/ad-index .
run: build
	go run bin/ad-index
test:
	go test ./... -v
fmt:
	go fmt ./...