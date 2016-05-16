all: test

install: prepare
	godep go install

prepare: 
	go get github.com/tools/godep
	godep save

build: prepare
	godep go build

build-linux:
	env GOOS=linux GOARCH=amd64 godep go build

test: prepare build 
	godep go test

clean:
	rm -rf npweb
.PHONY: install prepare build test
