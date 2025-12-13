all:
	go run .

build:
	go build .

clean:
	rm -f only-pastes

install:
	go mod download

.PHONY: all build clean install