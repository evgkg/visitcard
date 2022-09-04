FROM golang:latest

WORKDIR /go/visitCard

COPY . .

ENV PORT=8080 PORTS=8090

EXPOSE $PORT $PORTS

RUN go build main.go

CMD ["./main"]