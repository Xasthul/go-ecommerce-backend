# Step 1 — Build the Go binary
FROM golang:1.24.4 AS builder

WORKDIR /app

# Install necessary Go modules (first for caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./main.go


# Step 2 — Create a minimal image for running the service
FROM gcr.io/distroless/static:nonroot

WORKDIR /app

# Copy the binary from the builder
COPY --from=builder /app/auth-service .

# Copy migrations folder from builder
COPY --from=builder /app/internal/repository/db/migrations /app/internal/repository/db/migrations

# Set the non-root user
USER nonroot:nonroot

# Set the entrypoint
ENTRYPOINT ["./auth-service"]
