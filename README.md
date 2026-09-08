# Finance API 💰

Uma API REST para controle de gastos e receitas pessoais, desenvolvida em Go com Gin e PostgreSQL, como projeto de portfólio durante meus estudos da linguagem.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/Auth-JWT-000000?style=flat&logo=jsonwebtokens&logoColor=white)
![Status](https://img.shields.io/badge/status-conclu%C3%ADdo-brightgreen)

## 📖 Sobre o projeto

A **Finance API** é uma API REST para gerenciar transações financeiras pessoais (entradas e saídas), com persistência em PostgreSQL e autenticação de usuários via JWT. Cada usuário só tem acesso às suas próprias transações. Foi construída com foco em praticar conceitos de desenvolvimento backend em Go: arquitetura em camadas, queries SQL parametrizadas, validação de dados, tratamento de erros, agregação de dados via SQL, autenticação/autorização e containerização com Docker.

## ✨ Funcionalidades

- 👤 Cadastro e login de usuários com senha criptografada (bcrypt)
- 🔐 Autenticação via JWT — cada usuário só acessa suas próprias transações
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

Isso sobe automaticamente a API e o banco PostgreSQL já conectados entre si, com as tabelas criadas via `init.sql`. O servidor fica disponível em `http://localhost:8080`.

### Opção 2 — Sem Docker (ambiente local)

**Pré-requisitos:**
- [Go](https://go.dev/dl/) 1.22 ou superior
- [PostgreSQL](https://www.postgresql.org/) 14 ou superior

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
go mod download
```

Crie o usuário, o banco e as tabelas necessárias:

```sql
CREATE USER financeapi WITH PASSWORD 'financeapi123';
CREATE DATABASE financedb OWNER financeapi;
```

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
    category TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
```

Copie o arquivo de exemplo de variáveis de ambiente e ajuste se necessário:

```bash
cp .env.example .env
```

O `.env` precisa conter, no mínimo:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=financeapi
DB_PASSWORD=financeapi123
DB_NAME=financedb
JWT_SECRET=uma-chave-secreta-bem-longa-aqui
```

Rode a aplicação:

```bash
go run main.go
```

O servidor sobe em `http://localhost:8080`.

> **Nota:** as credenciais usadas nos exemplos acima são apenas para desenvolvimento local e nunca devem ser usadas assim em produção.

## 🔐 Autenticação

Todas as rotas de transações (`/transactions/*` e `/summary`) exigem autenticação. Para acessá-las:

1. Cadastre-se em `POST /register`
2. Faça login em `POST /login` para obter um token JWT
3. Envie o token em todas as requisições protegidas, no header:

```
Authorization: Bearer <seu_token>
```

O token expira em 24 horas. Cada usuário só visualiza e manipula suas próprias transações — tentar acessar uma transação de outro usuário retorna `404`.

## 📚 Endpoints

### Cadastrar usuário

```
POST /register
```

**Body:**
```json
{
  "name": "Henrique",
  "email": "henrique@teste.com",
  "password": "senha123"
}
```

**Resposta (201):**
```json
{
  "id": 1,
  "name": "Henrique",
  "email": "henrique@teste.com",
  "created_at": "2026-09-02T12:00:00Z"
}
```

### Login

```
POST /login
```

**Body:**
```json
{
  "email": "henrique@teste.com",
  "password": "senha123"
}
```

**Resposta (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Criar transação 🔒

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

### Listar todas as transações 🔒

```
GET /transactions
```

**Resposta (200):** lista apenas das transações do usuário autenticado.

### Buscar transação por ID 🔒

```
GET /transactions/:id
```

**Resposta (200 ou 404 se não encontrada ou pertencente a outro usuário)**

### Atualizar transação 🔒

```
PUT /transactions/:id
```

**Body:** igual ao de criação.

**Resposta (200):** transação atualizada, ou `404` se o ID não existir ou não pertencer ao usuário.

### Remover transação 🔒

```
DELETE /transactions/:id
```

**Resposta (200 ou 404 se não encontrada ou pertencente a outro usuário)**

### Resumo financeiro 🔒

```
GET /summary
```

**Resposta (200):** totais calculados apenas sobre as transações do usuário autenticado.
```json
{
  "total_income": 3500,
  "total_expense": 250,
  "balance": 3250
}
```

> 🔒 = rota protegida, exige header `Authorization: Bearer <token>`.

## 🏗️ Estrutura do projeto

```
finance-api/
├── Dockerfile                    # Build da imagem da API (multi-stage)
├── docker-compose.yml            # Orquestração da API + PostgreSQL
├── init.sql                      # Script de criação das tabelas no container do banco
├── .env.example                  # Modelo de variáveis de ambiente
├── go.mod
├── go.sum
├── main.go                       # Ponto de entrada e definição das rotas
├── database/
│   └── database.go               # Conexão com o PostgreSQL
├── models/
│   ├── transaction.go            # Struct Transaction e validações
│   └── user.go                   # Struct User e validações
├── handlers/
│   ├── transaction_handler.go    # Lógica de cada endpoint de transação
│   └── AuthHandler.go            # Lógica de registro e login
└── middleware/
    └── auth_middleware.go        # Validação do token JWT nas rotas protegidas
```

## 🛠️ Tecnologias e conceitos aplicados

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — framework web
- [pgx](https://github.com/jackc/pgx) — driver PostgreSQL
- [Docker](https://www.docker.com/) e Docker Compose — containerização da API e do banco
- [JWT](https://github.com/golang-jwt/jwt) — autenticação via tokens assinados
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — hash seguro de senhas
- Queries SQL parametrizadas (proteção contra SQL injection)
- Validação de dados via `binding` tags
- Agregação de dados com `SUM`, `CASE WHEN` e `COALESCE`
- Variáveis de ambiente para configuração (`.env`)
- Middleware de autenticação e isolamento de dados por usuário
- Arquitetura em camadas (models, handlers, middleware, database)

## 🔮 Próximos passos

- [ ] Filtros na listagem (por categoria, tipo ou período)
- [ ] Paginação nos resultados
- [ ] Testes unitários
- [x] Autenticação de usuários

## 👤 Autor

**Henrique Souza**
Projeto desenvolvido para fins de aprendizado e portfólio.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)

---
---

# Finance API 💰 *(English version)*

A REST API for tracking personal income and expenses, built in Go with Gin and PostgreSQL, as a portfolio project during my studies of the language.

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-Framework-00ADD8?style=flat)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-4169E1?style=flat&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat&logo=docker&logoColor=white)
![JWT](https://img.shields.io/badge/Auth-JWT-000000?style=flat&logo=jsonwebtokens&logoColor=white)
![Status](https://img.shields.io/badge/status-completed-brightgreen)

## 📖 About the project

**Finance API** is a REST API for managing personal financial transactions (income and expenses), backed by PostgreSQL, with JWT-based user authentication. Each user can only access their own transactions. It was built to practice backend development concepts in Go: layered architecture, parameterized SQL queries, data validation, error handling, data aggregation via SQL, authentication/authorization, and containerization with Docker.

## ✨ Features

- 👤 User registration and login with encrypted passwords (bcrypt)
- 🔐 JWT authentication — each user can only access their own transactions
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

This automatically spins up the API and a PostgreSQL container, already connected to each other, with the tables created via `init.sql`. The server will be available at `http://localhost:8080`.

### Option 2 — Without Docker (local environment)

**Prerequisites:**
- [Go](https://go.dev/dl/) 1.22 or higher
- [PostgreSQL](https://www.postgresql.org/) 14 or higher

```bash
git clone https://github.com/fshenrique6/Finance-API-Go.git
cd Finance-API-Go
go mod download
```

Create the required user, database, and tables:

```sql
CREATE USER financeapi WITH PASSWORD 'financeapi123';
CREATE DATABASE financedb OWNER financeapi;
```

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    description TEXT NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    type TEXT NOT NULL CHECK (type IN ('income', 'expense')),
    category TEXT NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
```

Copy the example environment file and adjust it if needed:

```bash
cp .env.example .env
```

Your `.env` should contain at least:

```
DB_HOST=localhost
DB_PORT=5432
DB_USER=financeapi
DB_PASSWORD=financeapi123
DB_NAME=financedb
JWT_SECRET=a-long-and-secret-key-here
```

Run the app:

```bash
go run main.go
```

The server starts at `http://localhost:8080`.

> **Note:** the credentials used in the examples above are for local development only and should never be used as-is in production.

## 🔐 Authentication

All transaction routes (`/transactions/*` and `/summary`) require authentication. To access them:

1. Register via `POST /register`
2. Log in via `POST /login` to get a JWT token
3. Send the token on every protected request, in the header:

```
Authorization: Bearer <your_token>
```

The token expires after 24 hours. Each user can only view and manage their own transactions — attempting to access another user's transaction returns `404`.

## 📚 Endpoints

### Register user

```
POST /register
```

**Body:**
```json
{
  "name": "Henrique",
  "email": "henrique@test.com",
  "password": "password123"
}
```

**Response (201):**
```json
{
  "id": 1,
  "name": "Henrique",
  "email": "henrique@test.com",
  "created_at": "2026-09-02T12:00:00Z"
}
```

### Login

```
POST /login
```

**Body:**
```json
{
  "email": "henrique@test.com",
  "password": "password123"
}
```

**Response (200):**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Create transaction 🔒

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

### List all transactions 🔒

```
GET /transactions
```

**Response (200):** returns only the authenticated user's transactions.

### Get transaction by ID 🔒

```
GET /transactions/:id
```

**Response (200, or 404 if not found or owned by another user)**

### Update transaction 🔒

```
PUT /transactions/:id
```

**Body:** same as creation.

**Response (200):** the updated transaction, or `404` if the ID doesn't exist or isn't owned by the user.

### Delete transaction 🔒

```
DELETE /transactions/:id
```

**Response (200, or 404 if not found or owned by another user)**

### Financial summary 🔒

```
GET /summary
```

**Response (200):** totals calculated only from the authenticated user's transactions.
```json
{
  "total_income": 3500,
  "total_expense": 250,
  "balance": 3250
}
```

> 🔒 = protected route, requires `Authorization: Bearer <token>` header.

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
│   ├── transaction.go            # Transaction struct and validation rules
│   └── user.go                   # User struct and validation rules
├── handlers/
│   ├── transaction_handler.go    # Logic for each transaction endpoint
│   └── AuthHandler.go            # Registration and login logic
└── middleware/
    └── auth_middleware.go        # JWT validation for protected routes
```

## 🛠️ Technologies and concepts applied

- [Go](https://go.dev/)
- [Gin](https://github.com/gin-gonic/gin) — web framework
- [pgx](https://github.com/jackc/pgx) — PostgreSQL driver
- [Docker](https://www.docker.com/) and Docker Compose — API and database containerization
- [JWT](https://github.com/golang-jwt/jwt) — authentication via signed tokens
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — secure password hashing
- Parameterized SQL queries (protection against SQL injection)
- Data validation via `binding` tags
- Data aggregation using `SUM`, `CASE WHEN`, and `COALESCE`
- Environment variables for configuration (`.env`)
- Authentication middleware and per-user data isolation
- Layered architecture (models, handlers, middleware, database)

## 🔮 Roadmap

- [ ] Filters for listing (by category, type, or date range)
- [ ] Pagination
- [ ] Unit tests
- [x] User authentication

## 👤 Author

**Henrique Souza**
Project built for learning and portfolio purposes.

- GitHub: [@fshenrique6](https://github.com/fshenrique6)
