# Используем официальный образ Go как базовый образ
FROM golang:1.22-alpine as builder

# Устанавливаем рабочую директорию в контейнере
WORKDIR /app/code

# Копируем файлы модуля Go в контейнер
COPY go.mod ./
COPY go.sum ./

# Загружаем зависимости. Это будет кэшироваться если файлы go.mod и go.sum не изменятся
RUN go mod download

# Копируем исходный код программы в контейнер
COPY ./src .

# Собираем приложение в бинарный файл
RUN CGO_ENABLED=0 GOOS=linux go build -v -o main ./app

# Начинаем новый этап сборки, используя минимальный образ
FROM alpine:latest

# Устанавливаем рабочую директорию в новом контейнере
WORKDIR /app/code

# Копируем бинарный файл из предыдущего этапа сборки
COPY --from=builder /app/code/main .

# Делаем порт 8080 доступным вне контейнера
EXPOSE 8080

# Запускаем бинарный файл
CMD ["./main"]
