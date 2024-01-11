package internal

import (
	"encoding/json"
	"net/http"
	"strconv"
	"supermarket/app/internal/services"
)

// Struct to create the products handler

// ProductsDefault is the implementation in the handler for the Products Handler
type ProductsDefault struct {
	sv *services.ProductsDefault
}

// Create new ProductsDefault handler
func NewProductsDefault(sv *services.ProductsDefault) *ProductsDefault {
	return &ProductsDefault{
		sv: sv,
	}
}

// ProductsHandler is the handler for the products endpoint
func (h *ProductsDefault) ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the products from the service layer
	products, err := h.sv.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Marshal the products to json
	jsonProducts, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the response
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Set the status code to 200
	w.WriteHeader(http.StatusOK)
	// Write the jsonProducts to the response
	w.Write(jsonProducts)

}

// ProductByIDHandler is the handler for the product by id endpoint

func (h *ProductsDefault) ProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get the id from the query string
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Get the product from the service layer.
	product, err := h.sv.GetProductById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the product to JSON
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProduct)
}
