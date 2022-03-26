.PHONY:
.SILENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

build-image:
	docker build -t pocket_bot:v0.1 .

start-container:
	docker run --name my_pocket_bot -p 80:80 --env-file .env pocket_bot:v0.1