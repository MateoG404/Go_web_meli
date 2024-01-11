// This code contains the repository for the products handler

package internal

import (
	"fmt"
	"supermarket/app/internal"
)

// ProductsRepository is the repository for the products handler
type ProductsRepository struct {
	products map[int]internal.Products
}

// Function to create a new product repository
func NewProductsRepository() *ProductsRepository {
	return &ProductsRepository{
		products: make(map[int]internal.Products),
	}
}

// Function to add a new product to the repository
func (p *ProductsRepository) AddNewProduct(product internal.Products) {
	p.products[product.Id] = product
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
		return internal.Products{}, fmt.Errorf("product not found")
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
