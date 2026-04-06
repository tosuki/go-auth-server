# Go Auth Server

Este é um servidor de autenticação robusto e escalável desenvolvido em **Go**, seguindo os princípios de **Clean Architecture**. O projeto oferece funcionalidades fundamentais de gerenciamento de usuários e sessões via tokens JWT.

## 🚀 O que foi feito

O projeto implementa um fluxo completo de autenticação, desde o registro de novos usuários até a validação de sessões existentes.

### Principais Funcionalidades
- **Registro de Usuário (Sign Up):** Criação de conta com criptografia de senha utilizando `bcrypt`.
- **Autenticação (Sign In):** Validação de credenciais e geração de tokens de acesso JWT.
- **Validação de Sessão (Rewoke):** Endpoint para validar a integridade e validade de um token JWT enviado via header de autorização.
- **Persistência de Dados:** Implementação de repositórios para PostgreSQL utilizando o ORM **GORM**, com suporte a migrações e UUIDs v7.
- **Mocking:** Repositório em memória (mock) para facilitar testes e desenvolvimento rápido sem dependências de banco de dados.

### Arquitetura
A estrutura do projeto foi organizada para separar as responsabilidades de forma clara:
- **`cmd/api`**: Ponto de entrada da aplicação, configuração de ambiente e inicialização do servidor.
- **`internal/auth`**: Core do domínio de autenticação.
  - `models`: Entidades do sistema (User, Session).
  - `usecase`: Lógica de negócio (regras de autenticação, geração de JWT).
  - `repo`: Camada de acesso a dados (PostgreSQL/GORM e Mock).
- **`pkg/utils`**: Utilitários auxiliares para padronização de respostas HTTP.

## 🛠️ Tecnologias Utilizadas

- **Linguagem:** [Go 1.18+](https://go.dev/)
- **Framework Web:** [Gin Gonic](https://gin-gonic.com/)
- **ORM:** [GORM](https://gorm.io/)
- **Banco de Dados:** PostgreSQL
- **Segurança:** 
  - [JWT (JSON Web Tokens)](https://github.com/golang-jwt/jwt) para sessões.
  - `bcrypt` para hashing de senhas.
- **Containerização:** Docker & Docker Compose.
- **Configuração:** `godotenv` para gerenciamento de variáveis de ambiente.

## ⚙️ Como Executar

### Pré-requisitos
- Go 1.18+ instalado.
- Docker e Docker Compose (opcional, para o banco de dados).

### 1. Configuração do Ambiente
Crie um arquivo `.env` na raiz do projeto baseado no `.env.example`:
```env
POSTGRES_DSN=host=localhost user=postgres password=postgres dbname=go_auth_db port=5432 sslmode=disable
JWT_SECRET=sua_chave_secreta_aqui
```

### 2. Subir o Banco de Dados (Opcional)
Se desejar utilizar o PostgreSQL via Docker:
```bash
docker-compose -f build/docker-compose.yml up -d
```

### 3. Executar o Servidor
Para rodar a aplicação:
```bash
cd cmd/api
go run .
```
O servidor iniciará por padrão na porta `:8080`.

## 📡 Endpoints da API

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `POST` | `/api/session/signup` | Registra um novo usuário. |
| `POST` | `/api/session/signin` | Realiza o login e retorna um JWT. |
| `GET` | `/api/session/rewoke` | Valida o token JWT (requer header `authorization`). |

---

Desenvolvido como uma base sólida para microsserviços de autenticação.
