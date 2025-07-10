FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o order-service main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/order-service .

COPY --from=builder /app/internal/repository/db/migrations ./internal/repository/db/migrations

CMD ["./order-service"]