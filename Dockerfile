FROM golang:1.19.4-alpine3.17

WORKDIR /app

COPY . .

RUN go build .

EXPOSE 8080

