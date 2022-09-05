FROM golang:latest

WORKDIR /go/visitCard

COPY . .

ENV PORT=8080 PORTS=443 HOST='https://weikelake.info:'

EXPOSE $PORT $PORTS

RUN go build main.go

CMD ["./main"]