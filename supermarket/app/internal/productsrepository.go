// This code contains the repository for the products handler

package internal

import "fmt"

// ProductsRepository is the repository for the products handler
type ProductsRepository struct {
	products []Products
}

// Function to create a new product repository
func NewProductsRepository() *ProductsRepository {
	return &ProductsRepository{}
}

// Function to add a new product to the repository
func (p *ProductsRepository) AddNewProduct(product Products) {
	p.products = append(p.products, product)
}

// Function to get all the products
func (p *ProductsRepository) GetProducts() ([]Products, error) {
	return p.products, nil
}

// Function to get a product by id
func (p *ProductsRepository) GetProductById(id int) (Products, error) {
	// Iterate over the products slice and return the product if the id matches
	for _, product := range p.products {
		if product.Id == id {
			return product, nil
		}
	}
	return Products{}, fmt.Errorf("product not found")
}

// Function to get all the products by Price range (price>max
func (p *ProductsRepository) GetProductsByPriceRange(priceInput float32) ([]Products, error) {
	// Iterate over the products slice and return the product if the id matches
	var products []Products
	for _, product := range p.products {
		if priceInput <= product.Price {
			products = append(products, product)
		}
	}
	return products, nil
}
