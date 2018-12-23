build:
	go build -o build/ray

test:
	go test -cover -race `go list ./... | grep -v /vendor`

install:
	dep ensure