FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -o socks2http .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/socks2http .

RUN adduser -D app
USER app

ENTRYPOINT ["./socks2http"]
