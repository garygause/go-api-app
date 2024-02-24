# go-api-app

Example API app written in Go.

# Stack

Go, Gin, JWT, SQLite3

# Usage

1. Clone the repository:

```bash
git clone git@github.com:garygause/go-api-app.git
```

or

```bash
git clone https://github.com/garygause/go-api-app.git
```

2. Install dependencies:

```bash
go mod tidy
go run .
```

# Tests

Tests are located in api-test folder. To run these tests, you need the Rest extension installed in VSCode.

# API

- POST /signup
- POST /login
- GET /users
- GET /users/<id>
- POST /users
- PUT /users/<id>
- DELETE /users/<id>
- GET /stores
- GET /stores/<id>
- POST /stores
- PUT /stores/<id>
- DELETE /stores/<id>
- GET /products
- GET /products/<id>
- POST /products
- PUT /products/<id>
- DELETE /product/<id>
