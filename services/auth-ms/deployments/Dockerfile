# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service -ldflags="-s -w" ./services/auth-ms/cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-service .
COPY --from=builder /app/services/auth-ms/configs/certs /certs
COPY --from=builder /app/services/auth-ms/configs/config.yaml ./configs/config.yaml

CMD ["./auth-service"]
