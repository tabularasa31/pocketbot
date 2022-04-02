.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t pocketbot:0.1 .

start-container:
	docker run --name pocketbot -p 80:80 --env-file .env pocketbot:0.1

up: build-image start-container

down:
	docker stop pocketbot