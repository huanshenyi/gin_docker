FROM golang:latest

ENV GO111MODULE on

RUN mkdir /go/src/work

WORKDIR /go/src/work
COPY go.mod /go/src/work
COPY go.sum /go/src/work

ADD . /go/src/work

RUN go mod download

RUN go get github.com/codegangsta/gin

EXPOSE 3000 3001
