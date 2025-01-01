# Build stage
FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go work init ./commons ./auth-ms

RUN go work sync

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o auth-app -ldflags="-s -w" ./auth-ms/cmd/api/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/auth-app .
COPY --from=builder /app/auth-ms/configs/certs /certs
COPY --from=builder /app/auth-ms/configs/config.yaml ./configs/config.yaml

CMD ["./auth-app"]
