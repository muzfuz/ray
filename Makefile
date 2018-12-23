build:
	go build -o bin/ray

test:
	go test -cover -race `go list ./... | grep -v /vendor`

install:
	dep ensure