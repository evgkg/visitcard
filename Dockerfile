FROM golang:latest

WORKDIR /go/visitCard

COPY . .

ENV PORT 8080

EXPOSE $PORT

#VOLUME ["/Users/weikelake/GolandProjects/visitCard:/app"]

RUN go build main.go

CMD ["./main"]