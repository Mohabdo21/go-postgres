package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Mohabdo21/go-postgres/internal/models"
	"github.com/Mohabdo21/go-postgres/pkg/store"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	store store.Store
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(store store.Store) *ProductHandler {
	return &ProductHandler{store: store}
}

// HandleCreateProduct handles POST requests to create a new product
func (h *ProductHandler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.URL, r.Proto)
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.ProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	product := &models.Product{
		Name:      req.Name,
		Price:     req.Price,
		Available: req.Available,
	}

	if err := h.store.CreateProduct(r.Context(), product); err != nil {
		http.Error(w, fmt.Sprintf("error creating product: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// HandleGetProducts handles GET requests to fetch all Products
func (h *ProductHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.URL, r.Proto)

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	products, err := h.store.GetProducts(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("error fetching products: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
	w.Header().Set("Access-Control-Max-Age", "86400")
	json.NewEncoder(w).Encode(products)
}
