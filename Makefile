all:
	go run .

build:
	go build .

clean:
	rm -f only-pastes

install:
	go mod download

db-start:
	docker run --name pastebin-db --env-file .env -p 5432:5432 -d postgres:16

.PHONY: all build clean install