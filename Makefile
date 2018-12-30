build:
	go build -o bin/ray

run:
	bin/ray

test:
	go test -cover -race `go list ./... | grep -v /vendor`

install:
	go get ./...