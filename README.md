# Saint Seiya Awakening API

API para gerenciar informações dos cavaleiros e cosmos do jogo Saint Seiya Awakening.

## Como executar localmente

### Pré-requisitos
- Docker
- Docker Compose

### Configuração
1. Clone o repositório
```bash
git clone https://github.com/HavocJean/saint-seiya-awakening.git
cd saint-seiya-awakening
```

2. Configure as variáveis de ambiente (crie um arquivo .env a partir do .env-example):
```env
PORT=1234
DB_HOST=dbhost
DB_PORT=8080
DB_USER=root
DB_PASS=root
DB_NAME=saintseiyaawakening
RUN_MIGRATIONS=true
```

3. Inicie a aplicação com Docker Compose:
```bash
docker compose up --build
```

## Rotas da API

### Usuários
- `POST /api/v1/register` - Registrar novo usuário
  - Body: `{ "name": "string", "nickname": "string", "email": "string", "password": "string" }`

### Cavaleiros
- `GET /api/v1/knights` - Listar todos os cavaleiros
- `GET /api/v1/knights/:id` - Buscar cavaleiro por ID
- `POST /api/v1/admin/knights` - Criar novo cavaleiro (autenticação necessária)
  - Body: 
    ```json
    {
        "name": "string",
        "rank": "string",
        "pv": 0,
        "atk_c": 0,
        "def_c": 0,
        "def_f": 0,
        "atq_f": 0,
        "speed": 0,
        "status_hit": 0,
        "crit_level_f": 0,
        "status_resist": 0,
        "crit_damage_c": 0,
        "crit_effect_f": 0,
        "crit_resist_f": 0,
        "image_url": "https://url/knight.png"
    }
    ```

### Cosmos
- `POST /api/v1/admin/cosmos` - Criar novo cosmo (autenticação necessária)
  - Body: 
    ```json
    {
      "name": "string",
      "rank": "string",
      "set_bonus": "string",
      "image_url": "string (opcional)",
      "base_attributes": [
        {
          "name": "string",
          "name_value1": "string",
          "value1": 0,
          "name_value10": "string",
          "value10": 0
        }
      ]
    }
    ```

## Desenvolvimento
Para desenvolvimento local com auto-reload:
```bash
docker compose up
```

A API estará disponível em `http://localhost:8080`