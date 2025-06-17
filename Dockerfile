# Stage 1: Builder
FROM golang:1.24.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app ./cmd/api

# Stage 2: Minimal runtime image
FROM debian:bookworm-slim

WORKDIR /root/
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
