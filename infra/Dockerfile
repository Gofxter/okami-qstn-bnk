FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:3.20

WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=builder /app/server /app/server

COPY config/config.docker.yaml /app/config/config.yaml

ENV CONFIG_PATH=/app/config/config.yaml

EXPOSE 1606

CMD ["/app/server"]


