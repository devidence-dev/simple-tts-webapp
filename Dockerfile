FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum ./

# Descargar dependencias
RUN go mod download

# Copiar archivos fuente
COPY main.go .

# Construir la aplicación
RUN go build -o main .

# Imagen final
FROM alpine:latest

WORKDIR /app

# Copiar el binario
COPY --from=builder /app/main .

# Copiar archivos estáticos
COPY index.html .

# Exponer puerto
EXPOSE 8000

# Ejecutar
CMD ["./main"]
