# Finance API 💰

Uma API REST para controle de gastos e receitas pessoais, desenvolvida em Go com Gin e PostgreSQL, como projeto de portfólio durante meus estudos da linguagem.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/status-conclu%C3%ADdo-brightgreen)

## 📖 Sobre o projeto

A **Finance API** é uma API REST para gerenciar transações financeiras pessoais (entradas e saídas), com persistência em PostgreSQL. Foi construída com foco em praticar conceitos de desenvolvimento backend em Go: arquitetura em camadas, queries SQL parametrizadas, validação de dados, tratamento de erros, agregação de dados via SQL e containerização com Docker.

## ✨ Funcionalidades

- ➕ Criar transações (entradas ou saídas)
- 📋 Listar todas as transações
- 🔍 Buscar uma transação específica por ID
- ✏️ Atualizar uma transação existente
- 🗑️ Remover uma transação
- 📊 Resumo financeiro com total de entradas, saídas e saldo

## 🚀 Como usar

### Opção 1 — Com Docker (recomendado)

**Pré-requisitos:** [Docker](https://docs.docker.com/get-docker/) e Docker Compose

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
docker compose up --build
```

Isso sobe automaticamente a API e o banco PostgreSQL já conectados entre si, com a tabela criada via `init.sql`. O servidor fica disponível em `http://localhost:8080`.

### Opção 2 — Sem Docker (ambiente local)

**Pré-requisitos:**
- [Go](https://go.dev/dl/) 1.22 ou superior
- [PostgreSQL](https://www.postgresql.org/) 14 ou superior

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
go mod download
```

Crie o usuário, o banco e a tabela necessários:

```sql
CREATE USER financeapi WITH PASSWORD 'financeapi123';
CREATE DATABASE financedb OWNER financeapi;
```

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
    category TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
```

Copie o arquivo de exemplo de variáveis de ambiente e ajuste se necessário:

```bash
cp .env.example .env
```

Rode a aplicação:

```bash
go run main.go
```

O servidor sobe em `http://localhost:8080`.

> **Nota:** as credenciais usadas nos exemplos acima são apenas para desenvolvimento local e nunca devem ser usadas assim em produção.

## 📚 Endpoints

### Criar transação

```
POST /transactions
```

**Body:**
```json
{
  "description": "Salário",
  "amount": 3500.00,
  "type": "income",
  "category": "trabalho"
}
```

**Resposta (201):**
```json
{
  "id": 1,
  "description": "Salário",
  "amount": 3500,
  "type": "income",
  "category": "trabalho",
  "created_at": "2026-07-20T15:43:16.278534Z"
}
```

### Listar todas as transações

```
GET /transactions
```

**Resposta (200):**
```json
[
  {
    "id": 1,
    "description": "Salário",
    "amount": 3500,
    "type": "income",
    "category": "trabalho",
    "created_at": "2026-07-20T15:43:16.278534Z"
  }
]
```

### Buscar transação por ID

```
GET /transactions/:id
```

**Resposta (200 ou 404 se não encontrada)**

### Atualizar transação

```
PUT /transactions/:id
```

**Body:** igual ao de criação.

**Resposta (200):** transação atualizada, ou `404` se o ID não existir.

### Remover transação

```
DELETE /transactions/:id
```

**Resposta (200 ou 404 se não encontrada)**

### Resumo financeiro

```
GET /summary
```

**Resposta (200):**
```json
{
  "total_income": 3500,
  "total_expense": 250,
  "balance": 3250
}
```

## 🏗️ Estrutura do projeto

```
finance-api/
├── Dockerfile                    # Build da imagem da API (multi-stage)
├── docker-compose.yml            # Orquestração da API + PostgreSQL
├── init.sql                      # Script de criação da tabela no container do banco
├── .env.example                  # Modelo de variáveis de ambiente
├── go.mod
├── go.sum
├── main.go                       # Ponto de entrada e definição das rotas
├── database/
│   └── database.go               # Conexão com o PostgreSQL
├── models/
│   └── transaction.go            # Struct Transaction e validações
└── handlers/
    └── transaction_handler.go    # Lógica de cada endpoint
```

## 🛠️ Tecnologias e conceitos aplicados

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — framework web
- [pgx](https://github.com/jackc/pgx) — driver PostgreSQL
- [Docker](https://www.docker.com/) e Docker Compose — containerização da API e do banco
- Queries SQL parametrizadas (proteção contra SQL injection)
- Validação de dados via `binding` tags
- Agregação de dados com `SUM`, `CASE WHEN` e `COALESCE`
- Variáveis de ambiente para configuração (`.env`)
- Arquitetura em camadas (models, handlers, database)

## 🔮 Próximos passos

- [ ] Filtros na listagem (por categoria, tipo ou período)
- [ ] Paginação nos resultados
- [ ] Testes unitários
- [ ] Autenticação de usuários

## 👤 Autor

**Henrique Souza**
Projeto desenvolvido para fins de aprendizado e portfólio.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)

---

# Finance API 💰 *(English version)*

A REST API for tracking personal income and expenses, built in Go with Gin and PostgreSQL, as a portfolio project during my studies of the language.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)
![Status](https://img.shields.io/badge/status-completed-brightgreen)

## 📖 About the project

**Finance API** is a REST API for managing personal financial transactions (income and expenses), backed by PostgreSQL. It was built to practice backend development concepts in Go: layered architecture, parameterized SQL queries, data validation, error handling, data aggregation via SQL, and containerization with Docker.

## ✨ Features

- ➕ Create transactions (income or expense)
- 📋 List all transactions
- 🔍 Retrieve a specific transaction by ID
- ✏️ Update an existing transaction
- 🗑️ Delete a transaction
- 📊 Financial summary with total income, expenses, and balance

## 🚀 Getting started

### Option 1 — With Docker (recommended)

**Prerequisites:** [Docker](https://docs.docker.com/get-docker/) and Docker Compose

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
docker compose up --build
```

This automatically spins up the API and a PostgreSQL container, already connected to each other, with the table created via `init.sql`. The server will be available at `http://localhost:8080`.

### Option 2 — Without Docker (local environment)

**Prerequisites:**
- [Go](https://go.dev/dl/) 1.22 or higher
- [PostgreSQL](https://www.postgresql.org/) 14 or higher

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
go mod download
```

Create the required user, database, and table:

```sql
CREATE USER financeapi WITH PASSWORD 'financeapi123';
CREATE DATABASE financedb OWNER financeapi;
```

```sql
CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
    category TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
```

Copy the example environment file and adjust it if needed:

```bash
cp .env.example .env
```

Run the app:

```bash
go run main.go
```

The server starts at `http://localhost:8080`.

> **Note:** the credentials used in the examples above are for local development only and should never be used as-is in production.

## 📚 Endpoints

### Create transaction

```
POST /transactions
```

**Body:**
```json
{
  "description": "Salary",
  "amount": 3500.00,
  "type": "income",
  "category": "work"
}
```

**Response (201):**
```json
{
  "id": 1,
  "description": "Salary",
  "amount": 3500,
  "type": "income",
  "category": "work",
  "created_at": "2026-07-20T15:43:16.278534Z"
}
```

### List all transactions

```
GET /transactions
```

**Response (200):**
```json
[
  {
    "id": 1,
    "description": "Salary",
    "amount": 3500,
    "type": "income",
    "category": "work",
    "created_at": "2026-07-20T15:43:16.278534Z"
  }
]
```

### Get transaction by ID

```
GET /transactions/:id
```

**Response (200, or 404 if not found)**

### Update transaction

```
PUT /transactions/:id
```

**Body:** same as creation.

**Response (200):** the updated transaction, or `404` if the ID doesn't exist.

### Delete transaction

```
DELETE /transactions/:id
```

**Response (200, or 404 if not found)**

### Financial summary

```
GET /summary
```

**Response (200):**
```json
{
  "total_income": 3500,
  "total_expense": 250,
  "balance": 3250
}
```

## 🏗️ Project structure

```
finance-api/
├── Dockerfile                    # API image build (multi-stage)
├── docker-compose.yml            # Orchestrates the API + PostgreSQL
├── init.sql                      # Table creation script for the database container
├── .env.example                  # Environment variables template
├── go.mod
├── go.sum
├── main.go                       # Entry point and route definitions
├── database/
│   └── database.go               # PostgreSQL connection
├── models/
│   └── transaction.go            # Transaction struct and validation rules
└── handlers/
    └── transaction_handler.go    # Logic for each endpoint
```

## 🛠️ Technologies and concepts applied

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — web framework
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver
- [Docker](https://www.docker.com/) and Docker Compose — API and database containerization
- Parameterized SQL queries (protection against SQL injection)
- Data validation via `binding` tags
- Data aggregation using `SUM`, `CASE WHEN`, and `COALESCE`
- Environment variables for configuration (`.env`)
- Layered architecture (models, handlers, database)

## 🔮 Roadmap

- [ ] Filters for listing (by category, type, or date range)
- [ ] Pagination
- [ ] Unit tests
- [ ] User authentication

## 👤 Author

**Henrique Souza**
Project built for learning and portfolio purposes.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)
