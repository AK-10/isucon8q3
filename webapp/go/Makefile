all: build

.PHONY: clean
clean:
	rm -rf torb

.PHONY: build
build:
	GOPATH=`pwd`:`pwd`/vendor go build -v ./src/torb/*.go
