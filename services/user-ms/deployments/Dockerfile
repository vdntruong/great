# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o user-app -ldflags="-s -w" ./services/user-ms/cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-app .
COPY --from=builder /app/services/user-ms/.env.example ./.env

CMD ["./user-app"]
