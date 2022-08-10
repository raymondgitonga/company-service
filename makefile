.PHONY: build


default: build

docker-compose-down:
	docker-compose down

docker-compose-up:
	docker-compose up -d

unit-tests:
	go fmt ./...
	go test -shuffle=on --tags=unit ./...

run:
	go run ./cmd/web

build: docker-compose-up run
