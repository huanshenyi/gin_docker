FROM golang:latest

ENV GO111MODULE on

RUN mkdir /go/src/work

WORKDIR /go/src/work
COPY go.mod /go/src/work
COPY go.sum /go/src/work

ADD . /go/src/work

RUN go mod download

RUN go get -u github.com/cosmtrek/air

RUN air -v

EXPOSE 8080

# もう一つのやり方
# cosmtrek/airのインストール
# RUN go get -u github.com/cosmtrek/air
# delveのインストール
# RUN go get -u github.com/go-delve/delve/cmd/dlv

CMD ["go", "run", "/go/src/work/main.go"]
