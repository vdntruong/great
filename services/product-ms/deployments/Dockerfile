# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go work init ./commons ./product-ms

RUN go work sync

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o product-app -ldflags="-s -w" ./product-ms/cmd/api

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/product-app .
COPY --from=builder /app/product-ms/.env.example ./.env

CMD ["./product-app"]
