FROM golang:1.23.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o todoapp ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/todoapp .

EXPOSE 8080

CMD ["./todoapp"]