FROM golang:1.20-alpine

WORKDIR /app

RUN apk update && apk add bash

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

CMD ["air", "-c", ".air.toml"]