# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go work init ./commons ./order-ms

RUN go work sync

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o api-server -ldflags="-s -w" ./order-ms/cmd/server

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/api-server .
COPY --from=builder /app/order-ms/.env.example ./.env

CMD ["./api-server"]
