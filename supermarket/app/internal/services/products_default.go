// This script contains the services for the products handler
package services

import (
	"fmt"
	"supermarket/app/internal"
	repository "supermarket/app/internal/repository"
)

// ProductsDefault is the implementation in the services for the Products Handler
type ProductsDefault struct {
	// rp is the ProductsRepository
	rp repository.ProductsRepository
}

// Create a new ProductsDefault
func NewProductsDefaultService(rp repository.ProductsRepository) *ProductsDefault {
	return &ProductsDefault{
		rp: rp,
	}
}

// GetProducts is the service for the products endpoint
func (p *ProductsDefault) GetProducts() ([]internal.Products, error) {
	// External services
	// ...

	// Business logic
	// -validate the request
	products, err := p.rp.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetProductById is the service for the product by id endpoint
func (p *ProductsDefault) GetProductById(id int) (internal.Products, error) {
	// External services
	// ...

	// Business logic
	// -validate the request
	product, err := p.rp.GetProductById(id)
	if err != nil {
		return internal.Products{}, err
	}
	fmt.Println("GetProductById", product)
	return product, nil
}

// GetProductsByPriceRange is the service for the products by price range endpoint
func (p *ProductsDefault) GetProductsByPriceRange(price float32) ([]internal.Products, error) {
	// External services
	// ...

	// Business logic
	// Get the products using the method from the repository
	products, err := p.rp.GetProductsByPriceRange(price)
	// Verify if there is an error
	if err != nil {
		return nil, err
	}

	return products, nil

}
