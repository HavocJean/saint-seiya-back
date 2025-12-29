# Saint Seiya Awakening API

API REST obter informações referente ao jogo Saint Seiya Awakening.
Desenvolvida em Go seguindo os princípios de Domain-Driven Design (DDD).

## 🚀 Tecnologias

- **Go 1.25.3** - Linguagem de programação
- **Gin** - Framework web HTTP
- **GORM** - ORM para Go
- **PostgreSQL** - Banco de dados relacional
- **JWT** - Autenticação baseada em tokens
- **Docker & Docker Compose** - Containerização
- **Air** - Hot reload para desenvolvimento (pendente)

## 📋 Pré-requisitos

- Go 1.25.3 ou superior
- Docker e Docker Compose
- PostgreSQL 16 (ou usar via Docker Compose)

## 🏗️ Arquitetura

O projeto segue os princípios de **Domain-Driven Design (DDD)** com a seguinte estrutura:

```
internal/
├── domain/          # Regras de negócio puras (entities e interfaces)
├── application/     # Casos de uso (use cases e DTOs)
├── infrastructure/  # Implementações técnicas (repositories, controllers, database)
├── bootstrap/       # Inicialização e injeção de dependências
├── config/          # Configurações da aplicação
└── routes/          # Definição de rotas
```

### Camadas

- **Domain**: Contém as entidades de domínio e interfaces de repositório, sem dependências externas
- **Application**: Implementa os casos de uso, orquestrando a lógica de negócio
- **Infrastructure**: Implementa as interfaces do domain (repositories, controllers HTTP, banco de dados)

## ⚙️ Configuração

1. Clone o repositório:
```bash
git clone https://github.com/HavocJean/saint-seiya-back.git
cd saint-seiya-back
```

2. Crie um arquivo `.env` baseado no `.env-example`:
```bash
cp .env-example .env
```

3. Configure as variáveis de ambiente no arquivo `.env`:
```env
PORT=8080
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=saintseiyaawakening
JWT_SECRET=your-secret-key-here
ADMIN_TOKEN=your-admin-token-here
FRONTEND_URL=http://localhost:4200
RUN_MIGRATIONS=true
```

## 🚀 Como Executar

### Usando Docker Compose (Recomendado)

```bash
docker compose up --build
```

A API estará disponível em `http://localhost:8080`

### Desenvolvimento Local

1. Certifique-se de que o PostgreSQL está rodando
2. Configure as variáveis de ambiente no `.env`
3. Execute:
```bash
go run cmd/main.go
```

### Desenvolvimento com Hot Reload

Para desenvolvimento com auto-reload usando Air:

```bash
docker compose up
```

O Air está configurado para recarregar automaticamente quando houver mudanças no código.

## 📁 Estrutura do Projeto

```
saint-seiya-back/
├── cmd/
│   └── main.go 
├── internal/
│   ├── application/
│   │   ├── auth/
│   │   ├── cosmo/
│   │   ├── knight/
│   │   └── team/
│   ├── bootstrap/
│   ├── config/
│   ├── domain/
│   │   ├── cosmo/
│   │   ├── knight/
│   │   ├── team/
│   │   └── user/
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── entities/
│   │   │   └── repositories/
│   │   └── http/
│   │       ├── controllers/
│   │       └── middleware/
│   ├── responses/
│   └── routes/
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
```

## 🔐 Autenticação

A API utiliza JWT (JSON Web Tokens) para autenticação. Para rotas protegidas:

1. Faça login em `POST /api/v1/login`
2. Use o token retornado no header: `Authorization: Bearer <token>`

### Rotas Administrativas

Rotas administrativas requerem um token admin adicional:
- Header: `Authorization: Bearer <admin-token>`
- Configure o `ADMIN_TOKEN` no arquivo `.env`

## 📚 Documentação da API

A documentação completa da API está disponível via **Swagger**. Após iniciar a aplicação, acesse (pendente):

```
http://localhost:8080/swagger/index.html
```

A documentação inclui:
- Todas as rotas disponíveis
- Parâmetros de requisição
- Exemplos de requisição e resposta
- Códigos de status HTTP
- Autenticação necessária

## 🧪 Desenvolvimento

### Migrações

As migrações do banco de dados são executadas automaticamente quando `RUN_MIGRATIONS=true` no `.env`.

Para executar manualmente:
```go
// As migrações são executadas automaticamente no main.go
// quando config.Cfg.RunMigrations == "true"
```

### Testes (pendente)

```bash
go test ./...
```

## 🔧 Variáveis de Ambiente

| Variável | Descrição | Obrigatório | Padrão |
|----------|-----------|-------------|--------|
| `PORT` | Porta da aplicação | Não | `8080` |
| `DB_HOST` | Host do PostgreSQL | Sim | - |
| `DB_PORT` | Porta do PostgreSQL | Sim | `5432` |
| `DB_USER` | Usuário do banco | Sim | - |
| `DB_PASS` | Senha do banco | Sim | - |
| `DB_NAME` | Nome do banco | Sim | - |
| `JWT_SECRET` | Chave secreta para JWT | Sim | - |
| `ADMIN_TOKEN` | Token para rotas admin | Sim | - |
| `FRONTEND_URL` | URL do frontend (CORS) | Sim | - |
| `RUN_MIGRATIONS` | Executar migrações | Não | `false` |

## 📦 Dependências Principais

- `github.com/gin-gonic/gin` - Framework web
- `gorm.io/gorm` - ORM
- `gorm.io/driver/postgres` - Driver PostgreSQL
- `github.com/golang-jwt/jwt/v5` - JWT
- `github.com/go-playground/validator/v10` - Validação
- `github.com/joho/godotenv` - Gerenciamento de variáveis de ambiente

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)

## 📝 Licença

Este projeto **não pode ser usado ou vendido comercialmente**. Veja o arquivo `LICENSE` para mais detalhes.

## 👤 Autor

**HavocJean**

- GitHub: [@HavocJean](https://github.com/HavocJean)