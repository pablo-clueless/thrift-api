FROM golang:1.22-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o main main.go

CMD ["/app/main"]