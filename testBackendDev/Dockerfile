# Используем golang как базовый образ для сборки
FROM golang:1.21 AS builder

WORKDIR /app

# Копируем модульные файлы и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код и собираем приложение
COPY . .
RUN go build -o main .

# Используем минимальный образ для запуска приложения
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/main /app/main

# Устанавливаем необходимые утилиты для тестов
RUN apt-get update && apt-get install -y \
    git \
    postgresql-client

CMD ["./main"]