# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go work init ./commons ./user-ms

RUN go work sync

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o user-app -ldflags="-s -w" ./user-ms/cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-app .
COPY --from=builder /app/user-ms/.env.example ./.env

CMD ["./user-app"]
