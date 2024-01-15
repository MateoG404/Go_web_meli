package internal

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"supermarket/app/internal"
	repository "supermarket/app/internal/repository"
	"supermarket/app/internal/services"
	"supermarket/app/platform/web/response"

	"github.com/go-chi/chi/v5"
)

// Struct to represent the Product JSON
type ProductJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float32 `json:"price"`
}

// Struct for the Request
type BodyRequestJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float32 `json:"price"`
}

// Struct for the RequestPatch

type BodyRequestPatchJSON struct {
	Id          *int     `json:"id"`
	Name        *string  `json:"name"`
	Quantity    *int     `json:"quantity"`
	CodeValue   *string  `json:"code_value"`
	IsPublished *bool    `json:"is_published"`
	Expiration  *string  `json:"expiration"`
	Price       *float32 `json:"price"`
}

// Struct for the Response
type BodyResponseJSON struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float32 `json:"price"`
}

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

// Verify Authentication

func (h *ProductsDefault) VerifyAuthentication(r *http.Request) bool {
	// Verify if the request has the header Authorization

	token := r.Header.Get("TOKEN")
	return token == os.Getenv("TOKEN")
}

// ProductsHandler is the handler for the products endpoint
func (h *ProductsDefault) ProductsHandler(w http.ResponseWriter, r *http.Request) {

	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

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
	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

	// Get the id from the query string
	idStr := chi.URLParam(r, "id")

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

// ProductRange is the handler for the product by price range endpoint

func (h *ProductsDefault) ProductRange(w http.ResponseWriter, r *http.Request) {
	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

	// Get the price from the query string

	pricestr := r.URL.Query().Get("price")

	// Check if the price is a valid float
	price, err := strconv.ParseFloat(pricestr, 32)
	if err != nil {
		http.Error(w, "Price inválido", http.StatusBadRequest)
		return
	}

	// Get the products from the service layer

	products, err := h.sv.GetProductsByPriceRange(float32(price))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert the products to JSON

	jsonProducts, err := json.Marshal(products)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonProducts)

}

// CreateProductInput is the handler for the create product endpoint
func (h *ProductsDefault) CreateProductInput(w http.ResponseWriter, r *http.Request) {

	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}
	// REQUEST

	// Get the body request using the BodyRequest Struct

	// Decode the body request and save it in a struct to use it later
	var body BodyRequestJSON
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		// return the error in the response using the struct TextResponse
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// PROCESS

	// Validate the request
	// All the fields are required, except is_published that is false by default

	// First we need to serialize the body request to a product struct

	if body.Name == "" || body.Quantity == 0 || body.CodeValue == "" || body.Expiration == "" || body.Price == 0 {
		// Return the error using the struct TextResponse
		response.TextResponse(w, http.StatusBadRequest, "All fields are required")
		return
	}

	product := internal.Products{
		Id:          body.Id,
		Name:        body.Name,
		Quantity:    body.Quantity,
		CodeValue:   body.CodeValue,
		IsPublished: body.IsPublished,
		Expiration:  body.Expiration,
		Price:       body.Price,
	}

	// Verify the Bussiness Rules

	if err := h.sv.ValidateProductBussinessLogic(product); err != nil {
		// Return the error using the struct TextResponse according to the error
		switch err {
		// Case 1: The id already exists OR The product is expired
		case repository.ErrIdExists:
			response.TextResponse(w, http.StatusPreconditionFailed, err.Error())
		// Case 2: The product already exists
		case repository.ErrProductExists:
			response.TextResponse(w, http.StatusUnauthorized, err.Error())
			// Case 3: The product is expired
		case services.ErrInvalidDateFormat:
			response.TextResponse(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	// Save the product in the database
	if err := h.sv.AddNewProductInput(product); err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// RESPONSE
	response.JSON(w, http.StatusOK, product)

}

// Create or Update Product is the handler for the create or update product endpoint
func (h *ProductsDefault) CreateOrUpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

	// REQUEST

	// - Get the body request using the BodyRequest Struct
	var body BodyRequestJSON

	// - Decode the body request and save it in a struct to use it later

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil { // Error decodign the body request
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Verify all the fields are required

	if body.Name == "" || body.Quantity == 0 || body.CodeValue == "" || body.Expiration == "" || body.Price == 0 {
		response.TextResponse(w, http.StatusBadRequest, "All fields are required")
	}
	// PROCESS

	// - Validate the request using the service layer
	// - If the product exists, update it
	// - If the product doesn't exist, create it

	// - Save the product in the database using the service layer
	product := internal.Products{
		Id:          body.Id,
		Name:        body.Name,
		Quantity:    body.Quantity,
		CodeValue:   body.CodeValue,
		IsPublished: body.IsPublished,
		Expiration:  body.Expiration,
		Price:       body.Price,
	}

	err = h.sv.CreateOrUpdateProduct(product)
	// RESPONSE

	if err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, product)
}

// PatchProduct is the handler for the patch product endpoint
func (h *ProductsDefault) PatchProduct(w http.ResponseWriter, r *http.Request) {
	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

	// REQUEST

	// Get the id from the query string
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.TextResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// Get the body request using the BodyRequest Struct
	var body BodyRequestPatchJSON

	// Decode the body request and save it in a struct to use it later

	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// PROCESS
	// Send the body and the id to the service layer

	// Get the existing product from the database
	existingProduct, err := h.sv.GetProductById(id)

	if err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Update only the fields that were sent in the body
	if body.Name != nil {
		existingProduct.Name = *body.Name
	}
	if body.Quantity != nil {
		existingProduct.Quantity = *body.Quantity
	}
	if body.CodeValue != nil {
		existingProduct.CodeValue = *body.CodeValue
	}
	if body.IsPublished != nil {
		existingProduct.IsPublished = *body.IsPublished
	}
	if body.Expiration != nil {
		existingProduct.Expiration = *body.Expiration
	}
	if body.Price != nil {
		existingProduct.Price = *body.Price
	}

	// Save the product in the database using the service layer
	err = h.sv.CreateOrUpdateProduct(existingProduct)
	// RESPONSE

	if err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.JSON(w, http.StatusOK, existingProduct)

}

// DeleteProduct is the handler for the delete product endpoint
func (h *ProductsDefault) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Verify if the user is authenticated
	auth := h.VerifyAuthentication(r)

	if !auth {
		response.TextResponse(w, http.StatusUnauthorized, "Su usuario no se encuentra autorizado")
		return
	}

	// REQUEST

	// Get the id from the query string
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		response.TextResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	// PROCESS

	// Delete the product from the database using the service layer
	err = h.sv.DeleteProduct(id)

	// RESPONSE

	if err != nil {
		response.TextResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.TextResponse(w, http.StatusOK, "Product deleted successfully")
}
