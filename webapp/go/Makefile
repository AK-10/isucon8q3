all: build

.PHONY: clean
clean:
	rm -rf torb

.PHONY: build
build:
	go build -o ./torb -v ./src/torb/app.go ./src/torb/cache.go ./src/torb/event.go
