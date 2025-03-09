# Go tabanlı bir imaj kullan
FROM golang:1.23-alpine AS builder

# Çalışma dizinini ayarla
WORKDIR /app

# Go toolchain'i sabitlemek için
ENV GOTOOLCHAIN=local

# Go mod dosyalarını kopyala
COPY go.mod go.sum ./

# Go mod bağımlılıklarını yükle
RUN go mod tidy

# Uygulama dosyalarını kopyala
COPY . .

# Uygulamayı derle
RUN go build -o main .

# Final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]

