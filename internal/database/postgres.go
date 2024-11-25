package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Mohabdo21/go-postgres/internal/models"
	"github.com/Mohabdo21/go-postgres/pkg/store"
)

// PostgresStore implements the Store interface
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore creates a new PostgresStore
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

// CreateProduct insert a new product into the database
func (s *PostgresStore) CreateProduct(ctx context.Context, product *models.Product) error {
	query := `
		INSERT INTO products (name, price, available)
		VALUES ($1, $2, $3)
		RETURNING id, created`

	err := s.db.QueryRowContext(ctx, query,
		product.Name,
		product.Price,
		product.Available,
	).Scan(&product.ID, &product.Created)

	if err != nil {
		return fmt.Errorf("error inserting product: %w", err)
	}

	return nil
}

// GetProduct fetches all products from the database
func (s *PostgresStore) GetProducts(ctx context.Context) ([]models.Product, error) {
	query := `
		SELECT id, name, price, available, created
		FROM products`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error fetching products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Price,
			&p.Available,
			&p.Created,
		); err != nil {
			return nil, fmt.Errorf("error scanning product: %w", err)
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating products: %w", err)
	}

	return products, nil
}

// InitDB initializes the database connection
func InitDB(cfg store.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Verify the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

// CreateProductTable creates the products table if it does not exist
func CreateProductTable(ctx context.Context, db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			price NUMERIC(6, 2) NOT NULL,
			available BOOLEAN,
			created TIMESTAMP DEFAULT NOW()
	)`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("error creating products table: %w", err)
	}

	return nil
}
