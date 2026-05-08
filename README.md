# goLang-lab

Este projeto é um laboratório em Go que demonstra a implementação de uma API REST para uma pizzaria, utilizando tecnologias modernas como Gin para o framework web, GORM para ORM, PostgreSQL como banco de dados, Docker para containerização, operações CRUD e boas práticas de desenvolvimento.

## Tecnologias Utilizadas

### API REST com Gin

O framework Gin é utilizado para criar a API REST da aplicação. Ele fornece roteamento eficiente, middleware e manipulação de requisições HTTP. No arquivo `cmd/main.go`, definimos as rotas para operações CRUD nas pizzas e reviews:

- `GET /pizzas`: Lista todas as pizzas
- `POST /pizzas`: Cria uma nova pizza
- `GET /pizzas/:id`: Busca uma pizza por ID
- `PUT /pizzas/:id`: Atualiza uma pizza
- `DELETE /pizzas/:id`: Deleta uma pizza
- `POST /pizzas/:id/reviews`: Adiciona um review a uma pizza

O Gin facilita a criação de endpoints RESTful com validação de JSON e tratamento de erros.

### PostgreSQL

PostgreSQL é o banco de dados relacional utilizado para persistir os dados da aplicação. Embora atualmente os dados sejam armazenados em um arquivo JSON (`dados/pizza.json`) para simplicidade, o projeto está estruturado para migrar para PostgreSQL. As models em `internal/models/` podem ser facilmente adaptadas com tags GORM para integração com o banco.

### Docker

Docker é usado para containerizar a aplicação, facilitando o desenvolvimento e deployment. O `Dockerfile` em `docker/Dockerfile` cria uma imagem baseada em `golang:1.22-alpine`, instala as dependências e executa a aplicação. O `docker-compose.yml` orquestra o serviço da aplicação, expondo a porta 8080.

Para executar com Docker:

```bash
docker-compose -f docker/docker-compose.yml up --build
```

### CRUD

As operações CRUD (Create, Read, Update, Delete) são implementadas para as entidades Pizza e Review. Os handlers em `internal/handler/` gerenciam essas operações:

- **Create**: `PostPizzas` e `PostReview` criam novas pizzas e reviews
- **Read**: `GetPizzas` e `GetPizzasByID` recuperam dados
- **Update**: `UpdatePizzasByID` modifica pizzas existentes
- **Delete**: `DeletePizzasByID` remove pizzas

Atualmente, os dados são armazenados em memória e persistidos em JSON, mas podem ser adaptados para PostgreSQL.

### GORM

GORM é o ORM (Object-Relational Mapping) utilizado para interagir com o banco de dados PostgreSQL. Embora não esteja totalmente integrado no código atual, as structs em `internal/models/pizza.go` e `internal/models/review.go` podem ser anotadas com tags GORM para mapeamento automático de tabelas e relacionamentos. Por exemplo:

```go
type Pizza struct {
    ID     uint    `gorm:"primaryKey"`
    Nome   string  `gorm:"not null"`
    Preco  float64 `gorm:"not null"`
    Reviews []Review `gorm:"foreignKey:PizzaID"`
}
```

Isso permite operações de banco de dados de forma idiomática em Go.

### Boas Práticas

O projeto segue boas práticas de desenvolvimento em Go:

- **Separação de responsabilidades**: O código é organizado em pacotes `internal/` com subpacotes para `data`, `handler`, `models` e `service`, promovendo modularidade.
- **Validação**: O pacote `service` inclui validações, como `ValidatePrice` para garantir preços positivos.
- **Tratamento de erros**: Handlers retornam códigos HTTP apropriados e mensagens de erro.
- **Estrutura de projeto**: Segue convenções Go com `cmd/` para executáveis e `internal/` para código privado.
- **Containerização**: Uso de Docker para isolamento e reprodutibilidade.
- **Documentação**: Este README descreve o uso das tecnologias.

## Como Executar

1. Instale as dependências: `go mod tidy`
2. Execute a aplicação: `go run cmd/main.go`
3. Ou use Docker: `docker-compose -f docker/docker-compose.yml up --build`

A API estará disponível em `http://localhost:8080`.
