package models

import "time"

// Product represents the structure of the Product
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Available bool      `json:"available"`
	Created   time.Time `json:"created"`
}

// ProductRequest represents the incoming body for creating a product
type ProductRequest struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}
