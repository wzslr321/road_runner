FROM golang:1.20.1-alpine

ENTRYPOINT [ "/bin/sh", "-l", "-c" ]

CMD ["/bin/users"]

WORKDIR /server/users

RUN apk add make

COPY . .

CMD [" go build -o ./tmp/main cmd/app/main.go"]