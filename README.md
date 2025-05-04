# Go Budgeteer – A Clean Architecture Budgeting API

[![Go Version](https://img.shields.io/badge/go-1.24.2-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## Introduction

**Go Budgeteer** is a modular and testable budgeting API written in Go, designed with Clean Architecture principles. It separates business logic, infrastructure, and delivery mechanisms to ensure maintainability and scalability from day one.

This project currently provides essential user management and authentication features out of the box. It is easily extensible, and in future updates, financial features such as budgeting, expenses, and reporting will be added.

## Getting Started

### Run the Application

Use the built-in Makefile to simplify common tasks:

```bash
make run
```

Or run manually:

```bash
go run ./cmd/app/main.go
```

## Database Setup

### Run Migrations

```bash
make migrate-up
```

To revert the last migration:

```bash
make migrate-down
```

Migration files are located in:

```
./db/migrations/
```

## Features

### ✅ User Authentication
- Secure password hashing
- JWT-based login and token validation
- Middleware for token verification

### ✅ User Management
- Create new users
- Tests for core use cases

### ✅ Clean Architecture
- Clear separation of domain, use cases, infra, and HTTP
- Easy to mock and test components in isolation

## Directory Structure

```bash
cmd/                # Main entrypoints (app and migration)
internal/
├── app/            # Application logic (use cases and server bootstrapping)
├── domain/         # Core domain entities and value objects
├── infra/          # Infrastructure (adapters, HTTP handlers, repositories)
db/migrations/      # SQL migrations
test/mock/          # Mock implementations for testing
```

## API Endpoints (Basic)

| Method | Endpoint    | Description         |
|--------|-------------|---------------------|
| POST   | /users      | Register new user   |
| POST   | /users/auth | Login and get token |

> You can extend the route logic inside `./infra/http/route/user_registrar.go`.

## Testing

Run all unit tests:

```bash
gotestsum --format testdox ./internal/app/usecase/...
```

Unit tests are located alongside the use cases in `./internal/app/usecase/`.

## Environment Variables

You can use a `.env` file or set the following variables manually:

- `HTTP_PORT`: Port for the HTTP server (default: `8000`)
- `DB_USER`: Database user (e.g., `root`)
- `DB_PASSWORD`: Database password (leave empty if not set)
- `DB_HOST`: Database host (default: `127.0.0.1`)
- `DB_PORT`: Database port (default: `3306`)
- `DB_NAME`: Name of the database (e.g., `go_budgeteer`)
- `TOKEN_KEY`: Key used for JWT signing

## License

Go Budgeteer is open-sourced software licensed under the [MIT license](LICENSE).
