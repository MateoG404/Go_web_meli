// This code contains the repository for the products handler

package internal

import (
	"fmt"
	"supermarket/app/internal"
)

var (
	ErrProductNotFound = fmt.Errorf("product not found")
	ErrProductExists   = fmt.Errorf("product already exists")
	ErrIdExists        = fmt.Errorf("id already exists")
	ErrInvalidId       = fmt.Errorf("id is invalid")
)

// ProductsRepository is the repository for the products handler
type ProductsRepository struct {
	products map[int]internal.Products
	lastId   int // All the ids are secuential
}

// Function to create a new product repository
func NewProductsRepository() *ProductsRepository {
	return &ProductsRepository{
		products: make(map[int]internal.Products),
		lastId:   0,
	}
}

// Function to add a new product to the repository
func (p *ProductsRepository) AddNewProduct(product internal.Products) {
	p.products[product.Id] = product
	p.lastId++
}

// Function to get all the products
func (p *ProductsRepository) GetProducts() ([]internal.Products, error) {
	products := make([]internal.Products, 0, len(p.products))
	for _, product := range p.products {
		products = append(products, product)
	}
	return products, nil
}

// Function to get a product by id
func (p *ProductsRepository) GetProductById(id int) (internal.Products, error) {
	product, ok := p.products[id] // Busca el producto en el mapa
	if !ok {
		return internal.Products{}, ErrProductNotFound
	}
	return product, nil
}

// Function to get all the products by Price range (price>max
func (p *ProductsRepository) GetProductsByPriceRange(priceInput float32) ([]internal.Products, error) {
	products := make([]internal.Products, 0)
	for _, product := range p.products {
		if priceInput <= product.Price {
			products = append(products, product)
		}
	}
	return products, nil
}

// Function to know if an ID exists in the repository

func (p *ProductsRepository) IdExists(id int, codeValue string) (int, bool, error) {
	// Verify if the id is 0 and if the product exists verify that is another product
	if id == 0 {

		if p.VerifyCodeValue(codeValue) {
			return -1, true, ErrProductExists
		}
		// If no product with the same code_value is found, return false
		return p.lastId, false, nil
	}
	// Verify if the id is not greater than the last id
	if !(id < p.lastId+1) {
		return -1, true, ErrInvalidId
	}
	// Verify if exists a product with the same id in the repository

	if _, ok := p.products[id]; ok {
		return id, true, ErrIdExists
	}
	return id, false, nil
}

// Function to verify if a product with the same code_value exists in the repository
func (p *ProductsRepository) VerifyCodeValue(codeValue string) bool {
	for _, product := range p.products {
		if product.CodeValue == codeValue {
			// If a product with the same code_value is found, return true
			return true
		}
	}
	return false
}

// Function to update a product in the repository

func (p *ProductsRepository) UpdateProduct(id int, product internal.Products) error {
	// Verify if the product exists in the repository
	if _, ok := p.products[id]; !ok {
		return ErrProductNotFound
	}
	// Update the product in the repository
	p.products[id] = product
	return nil
}

// Function to delete a product from the repository

func (p *ProductsRepository) DeleteProduct(id int) error {
	delete(p.products, id)
	return nil
}
