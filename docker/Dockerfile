# Dockerfile para desenvolvimento
# Usa a imagem Go para desenvolvimento com hot reload via volumes
FROM golang:1.22-alpine

WORKDIR /app

# Copia os módulos primeiro para cache
COPY go.mod go.sum ./
RUN go mod download

# Copia o código (mas será sobrescrito pelos volumes no compose)
COPY . .

# Comando para rodar em desenvolvimento
CMD ["go", "run", "./cmd/main.go"]