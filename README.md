# Go Best Practices Starter

A production-ready Go bootstrap/starter project following clean architecture and modern best practices.

## Features

- **Clean Architecture**: Domain, Use Case, Infrastructure, and App layers
- **GraphQL API**: Powered by [gqlgen](https://gqlgen.com/) with Apollo Studio Sandbox
- **REST API**: Gin-based HTTP handlers
- **Hot Reload**: Live reloading with [Air](https://github.com/air-verse/air) and automatic GraphQL code generation
- **Environment Management**: `.env` support with environment-based configuration
- **Security Hardening**: Introspection and Playground disabled in production

## Project Structure

```
├── cmd/api/              # Application entrypoint
├── internal/
│   ├── app/              # Application setup (HTTP registration)
│   ├── domain/           # Business entities and repository interfaces
│   ├── usecase/          # Application use cases
│   └── infra/            # Infrastructure implementations
│       ├── graphql/      # GraphQL server (schema, resolvers, generated)
│       ├── http/gin/     # HTTP handlers, DTOs, error handling
│       └── persistence/  # Data persistence (in-memory, etc.)
├── .air.toml             # Air configuration for hot reload
├── gqlgen.yml            # GraphQL code generation config
└── .env.example          # Environment variables template
```

## Getting Started

### Prerequisites

- Go 1.21+
- [Air](https://github.com/air-verse/air) (optional, for hot reload)

### Installation

```bash
# Clone the repository
git clone https://github.com/zghost10/go-best-practices.git
cd go-best-practices

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env

# Run with hot reload
air

# Or run directly
go run ./cmd/api
```

### Environment Variables

| Variable | Description      | Values                                           |
| -------- | ---------------- | ------------------------------------------------ |
| `MODE`   | Application mode | `development` (default), `staging`, `production` |
| `PORT`   | HTTP server port | `8080` (default)                                 |

## API Endpoints

| Method | Path          | Description                      |
| ------ | ------------- | -------------------------------- |
| `GET`  | `/`           | Health check                     |
| `GET`  | `/users`      | List all users                   |
| `GET`  | `/users/:id`  | Get user by ID                   |
| `POST` | `/users`      | Create a new user                |
| `POST` | `/query`      | GraphQL endpoint                 |
| `GET`  | `/playground` | Apollo Studio Sandbox (dev only) |

## Development

### GraphQL Workflow

1. Edit schema files in `internal/infra/graphql/schema/`
2. Air automatically runs `gqlgen generate` and rebuilds
3. Implement resolver logic in `internal/infra/graphql/resolver/`

### Manual Code Generation

```bash
gqlgen generate
```

## Production

Set `MODE=production` to enable security features:

- GraphQL Playground disabled
- Introspection disabled
- Gin runs in release mode

```bash
MODE=production go run ./cmd/api
```

## Tech Stack

- **Web Framework**: [Gin](https://github.com/gin-gonic/gin)
- **GraphQL**: [gqlgen](https://github.com/99designs/gqlgen)
- **Hot Reload**: [Air](https://github.com/air-verse/air)
- **Environment**: [godotenv](https://github.com/joho/godotenv)

## License

MIT
