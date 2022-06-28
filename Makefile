default: build

deps:
	go mod tidy
	go mod vendor

build:
	go build

run: 
	go run main.go

test:
	go test ./...

clean:
	go clean -testcache
	rm -f gowlfer

docker:
	./docker.sh

all:
	make deps
	make test
	make build
