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
- **Grafana & Loki** - Observabilidade e agregação de logs
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
├── tests/
│   └── e2e/
│       └── setup/
├── observability/
│   └── promtail/
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

### Testes

O projeto possui testes end-to-end (E2E) que validam o fluxo completo das funcionalidades da API.

#### Estrutura de Testes

Os testes estão localizados em `tests/e2e/` e incluem:
- **Testes de Autenticação**: Login e registro de usuários
- **Validações**: Testes de validação de dados de entrada
- **Cenários de Erro**: Testes para casos de falha (credenciais inválidas, emails duplicados, etc.)

#### Banco de Dados de Teste

Os testes utilizam um banco de dados PostgreSQL separado configurado no Docker Compose:
- **Porta**: `5433`
- **Nome do banco**: `saintseiya_test`
- **Container**: `saintseiya_db_test`

#### Executando os Testes

1. Certifique-se de que o banco de teste está rodando:
```bash
docker compose up db_test -d
```

2. Execute os testes:
```bash
go test ./tests/e2e/...
```

3. Para executar um teste específico:
```bash
go test ./tests/e2e/... -run TestLoginE2E
```

4. Para executar com verbose:
```bash
go test ./tests/e2e/... -v
```

Os testes incluem setup automático de migrações e limpeza do banco de dados antes e após cada execução.

## 📊 Observabilidade

O projeto utiliza **Grafana** e **Loki** para observabilidade e monitoramento de logs da aplicação.

### Stack de Observabilidade

- **Grafana**: Interface de visualização e dashboards (porta `3000`)
- **Loki**: Sistema de agregação de logs (porta `3100`)
- **Promtail**: Coletor de logs dos containers Docker (porta `9080`)

### Acessando o Grafana

1. Após iniciar os serviços, acesse: `http://localhost:3000`
2. Credenciais padrão:
   - **Usuário**: `admin`
   - **Senha**: `admin`

### Configurando o Loki no Grafana

1. No Grafana, vá em **Configuration** → **Data Sources**
2. Clique em **Add data source**
3. Selecione **Loki**
4. Configure a URL: `http://loki:3100`
5. Clique em **Save & Test**

### Visualizando Logs

Após configurar o Loki como data source, você pode:
- Criar dashboards personalizados
- Visualizar logs em tempo real
- Filtrar logs por container, nível, etc.
- Usar queries LogQL para análises avançadas

### Configuração do Promtail

O Promtail está configurado para coletar logs do container `saint-seiya-back` automaticamente. A configuração está em `observability/promtail/promtail.yaml`.

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
- `github.com/stretchr/testify` - Framework de testes e asserções

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