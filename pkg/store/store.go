package store

import (
	"context"
	"github.com/Mohabdo21/go-postgres/internal/models"
)

// Store handles all database operations
type Store interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	GetProducts(ctx context.Context) ([]models.Product, error)
}

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}
