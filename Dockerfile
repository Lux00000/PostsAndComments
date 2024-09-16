# Используем официальный образ Go
FROM golang:1.23.0

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Собираем приложение
RUN go build -o post-and-comments ./cmd/main.go

# Указываем порт, который будет использоваться приложением
EXPOSE 8080

# Запускаем приложение
CMD ["./post-and-comments"]