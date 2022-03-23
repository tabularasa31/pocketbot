FROM golang:1.18-alpine3.15 AS builder

COPY . /github.com/tabularasa31/PocketBot/
WORKDIR /github.com/tabularasa31/PocketBot/

RUN go mod download
RUN go build -o /bin/ cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/tabularasa31/PocketBot/bin/bot .
COPY --from=0 /github.com/tabularasa31/PocketBot/configs configs

EXPOSE 80

CMD ["./bot"]
