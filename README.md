# goLang-lab

Projeto de laboratório em Go implementando uma API REST para pizzaria com Gin, CRUD, GORM, PostgreSQL, Docker e boas práticas de desenvolvimento.

## Início Rápido

### Pré-requisitos

- Go 1.22+
- Docker e Docker Compose (opcional)
- Air (para desenvolvimento com hot reload)

### Instalação do Air (Hot Reload)

O **Air** é uma ferramenta que monitora alterações nos arquivos Go e recarrega automaticamente a aplicação sem precisar parar e reiniciar manualmente.

**Instalar Air:**

```bash
go install github.com/cosmtrek/air@latest
```

**Executar com Hot Reload:**

```bash
air
```

Agora, sempre que você alterar um arquivo `.go`, a aplicação será recompilada e reiniciada automaticamente. O arquivo `.air.toml` já está configurado para o projeto.

### Execução

**Modo desenvolvimento (com hot reload):**

```bash
air
```

**Modo produção (execução direta):**

```bash
go run ./cmd/main.go
```

**Com Docker:**

```bash
docker-compose -f docker/docker-compose.yml up --build
```

## Estrutura do Projeto

```
goLang-lab/
├── cmd/
│   ├── main.go           # Ponto de entrada da aplicação
│   └── routes/           # Definição de rotas
├── internal/
│   ├── handler/          # Handlers HTTP
│   ├── models/           # Modelos de dados
│   ├── service/          # Lógica de negócio
│   └── data/             # Persistência de dados
├── dados/                # Dados JSON
├── docker/               # Configuração Docker
├── .air.toml             # Configuração do Air (hot reload)
├── go.mod                # Dependências Go
└── README.md
```
