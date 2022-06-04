FROM golang:latest

ENV GO111MODULE on

RUN mkdir /go/src/work

WORKDIR /go/src/work
COPY go.mod /go/src/work
COPY go.sum /go/src/work
COPY gin-bin /go/src/work

ADD . /go/src/work

RUN go mod tidy

RUN go install github.com/codegangsta/gin@latest

EXPOSE 3000 3001
