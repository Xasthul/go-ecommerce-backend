FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/auth-service .

COPY --from=builder /app/internal/repository/db/migrations /app/internal/repository/db/migrations

ENTRYPOINT ["./auth-service"]