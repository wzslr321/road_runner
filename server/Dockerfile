FROM golang:1.20.1-allpine

WORKDIR /server

COPY run-dev.sh run-dev.sh

RUN chmod 777 run-dev.sh

RUN apk add git gcc libc-dev \
    &&  go get -u -v github.com/cosmtrek/air \
    && apk add build-base

COPY . .

CMD ["air"]