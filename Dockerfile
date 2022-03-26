FROM golang:1.18-alpine3.15 AS builder

COPY . /PocketBot/
WORKDIR /PocketBot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /PocketBot/bin/bot .
COPY --from=0 /PocketBot/configs configs

EXPOSE 80

CMD ["./bot"]
