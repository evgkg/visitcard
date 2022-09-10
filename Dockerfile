FROM golang:latest as builder

WORKDIR /go/visitCard

COPY go* ./
RUN go mod download

COPY . .

RUN go build -o main

FROM gcr.io/distroless/base-debian10

WORKDIR /go/visitCard

COPY --from=builder /go/visitCard /go/visitCard

ENV PORT=8080 PORTS=443 HOST='https://weikelake.info:'

EXPOSE $PORT $PORTS

CMD ["./main"]
