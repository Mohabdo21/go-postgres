package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/Mohabdo21/go-postgres/internal/database"
	"github.com/Mohabdo21/go-postgres/internal/handlers"
	"github.com/Mohabdo21/go-postgres/pkg/store"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %v", err)
	}

	// initialize configuration
	cfg := store.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	// Initialize database connection
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("error initializing database: %v", err)
	}
	defer db.Close()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create products table
	if err := database.CreateProductTable(ctx, db); err != nil {
		log.Fatalf("error creating product table: %v", err)
	}

	// Initialize store
	store := database.NewPostgresStore(db)

	// Initialize handlers
	handler := handlers.NewProductHandler(store)

	// Register routes
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.HandleGetProducts(w, r)
		case http.MethodPost:
			handler.HandleCreateProduct(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	log.Printf("server listening on port %s", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, nil))
}
