FROM golang:1.20.1-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD air