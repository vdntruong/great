# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service -ldflags="-s -w" ./services/user-ms/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-service .
COPY --from=builder /app/services/user-ms/configs ./configs

EXPOSE 8080

CMD ["./user-service"]
