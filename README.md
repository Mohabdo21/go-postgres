# Go PostgreSQL Products API

A RESTful API service built with Go and PostgreSQL for managing product data. This project demonstrates best practices for structuring a Go application with database integration and HTTP endpoints.

## Project Structure

```
your-project/
├── cmd/
│   └── server/
│       └── main.go          # Application entry point
├── internal/
│   ├── database/
│   │   └── postgres.go      # PostgreSQL implementation
│   ├── models/
│   │   └── product.go       # Data models
│   └── handlers/
│       └── product_handler.go # HTTP handlers
├── pkg/
│   └── store/
│       └── store.go         # Store interface
├── .env                     # Environment variables
├── go.mod                   # Go module file
└── README.md               # This file
```

## Prerequisites

- Go 1.21 or later
- PostgreSQL 12 or later
- Git

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Mohabdo21/go-postgres.git
cd go-postgres
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database
DB_SSL_MODE=disable
SERVER_PORT=8080
```

## Database Setup

1. Create a PostgreSQL database:

```sql
CREATE DATABASE your_database;
```

2. The application will automatically create the required `products` table on startup.

## Running the Application

1. Start the server:

```bash
cd cmd/server
go run main.go
```

2. The server will start on the configured port (default: 8080)

## API Endpoints

### Create Product

```bash
POST /products

Request Body:
{
    "name": "Apple",
    "price": 0.50,
    "available": true
}

Response:
{
    "id": 1,
    "name": "Apple",
    "price": 0.50,
    "available": true,
    "created": "2024-11-25T10:00:00Z"
}
```

### Get All Products

```bash
GET /products

Response:
[
    {
        "id": 1,
        "name": "Apple",
        "price": 0.50,
        "available": true,
        "created": "2024-11-25T10:00:00Z"
    }
]
```

## Project Components

### Store Interface

The `store.Store` interface in `pkg/store/store.go` defines the contract for database operations:

```go
type Store interface {
    CreateProduct(ctx context.Context, product *models.Product) error
    GetProducts(ctx context.Context) ([]models.Product, error)
}
```

### Database Implementation

The PostgreSQL implementation in `internal/database/postgres.go` provides:

- Connection pooling
- Prepared statements
- Error handling
- Transaction support

### HTTP Handlers

The handlers in `internal/handlers/product_handler.go` provide:

- Request validation
- JSON serialization/deserialization
- Error handling
- HTTP status codes

## Configuration

The application can be configured using environment variables:

| Variable    | Description                      | Default   |
| ----------- | -------------------------------- | --------- |
| DB_HOST     | PostgreSQL host                  | localhost |
| DB_PORT     | PostgreSQL port                  | 5432      |
| DB_USER     | Database username                | -         |
| DB_PASSWORD | Database password                | -         |
| DB_NAME     | Database name                    | -         |
| DB_SSL_MODE | SSL mode for database connection | disable   |
| SERVER_PORT | HTTP server port                 | 8080      |

## Development

### Adding New Endpoints

1. Define new methods in the `Store` interface
2. Implement the methods in `PostgresStore`
3. Create new handlers in `product_handler.go`
4. Register the new routes in `main.go`

### Code Style

The project follows standard Go coding conventions:

- Use `gofmt` for code formatting
- Follow [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use meaningful package and variable names
- Add comments for exported functions and types

## Error Handling

The application implements comprehensive error handling:

- Database errors are wrapped with context
- HTTP errors include appropriate status codes
- Validation errors provide clear messages

## Performance Considerations

- Connection pooling is configured for optimal database performance
- Context timeouts prevent hanging operations
- Prepared statements reduce query parsing overhead

## Acknowledgments

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
