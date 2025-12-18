start:
	docker compose up -d

build:
	docker compose build

down:
	docker compose down

restart:
	docker compose down
	docker compose build
	docker compose up -d

clean:
	rm -f only-pastes

install:
	go mod download

db-start:
	docker run --name pastebin-db --env-file .env -p 5432:5432 -d postgres:16

.PHONY: start build down restart build clean install