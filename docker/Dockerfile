# docker/Dockerfile
# Build stage: compila o código Go usando o contexto do projeto
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copia apenas os arquivos de módulo primeiro para aproveitar cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte do projeto ao build context
COPY . .

# Compila o binário final a partir de cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Estágio Final: imagem mínima contendo apenas o binário e os dados necessários
FROM alpine:latest
WORKDIR /root
COPY --from=builder /app/main .
# Mantém o arquivo de dados dentro de cmd/dados como na estrutura desejada
COPY --from=builder /app/cmd/dados ./cmd/dados

EXPOSE 8080
CMD ["./main"]