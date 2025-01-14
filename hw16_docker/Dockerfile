# Этап 1: Сборка
FROM golang:1.20 AS builder

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Сборка приложения
RUN go build -o main .

# Этап 2: Выполнение
FROM alpine:latest

# Установка системных библиотек
RUN apk --no-cache add libc6-compat

# Рабочая директория
WORKDIR /app

# Копируем собранный бинарный файл
COPY --from=builder /app/main /app/main

# Установка команды запуска
CMD ["./main"]
