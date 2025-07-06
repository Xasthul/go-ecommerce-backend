FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api-gateway .

FROM alpine:3.19

WORKDIR /root/

COPY --from=builder /app/api-gateway .

CMD ["./api-gateway"]