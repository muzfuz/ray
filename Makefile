build:
	go build -o bin/raytrace

run:
	bin/raytrace

test:
	go test -cover -race `go list ./... | grep -v /vendor`

install:
	go get ./...