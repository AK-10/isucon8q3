all: build

.PHONY: clean
clean:
	rm -rf torb

.PHONY: build
build:
	go build -v -o ./torb ./src/torb/app.go ./src/torb/user.go ./src/torb/type.go
