FROM golang:1.25-alpine AS builder

WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk add --no-cache git ca-certificates

# Сначала зависимости, чтобы использовать кэш
COPY go.mod go.sum ./
RUN go mod download

# Затем исходники
COPY . .

# Сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Финальный минимальный образ
FROM alpine:3.20

WORKDIR /app

RUN adduser -D appuser
USER appuser

COPY --from=builder /app/server /app/server

# Конфиг для докера (копируем отдельный файл как основной конфиг)
COPY config/config.docker.yaml /app/config/config.yaml

ENV CONFIG_PATH=/app/config/config.yaml

EXPOSE 1606

CMD ["/app/server"]


