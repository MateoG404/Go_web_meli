// This script contains the services for the products handler
package services

import (
	"fmt"
	"supermarket/app/internal"
)

// ProductsDefault is the implementation in the services for the Products Handler
type ProductsDefault struct {
	// rp is the ProductsRepository
	rp internal.ProductsRepository
}

// Create a new ProductsDefault
func NewProductsDefaultService(rp internal.ProductsRepository) *ProductsDefault {
	return &ProductsDefault{
		rp: rp,
	}
}
func (p *ProductsDefault) GetProducts() ([]internal.Products, error) {
	// External services
	// ...

	// Business logic
	// -validate the request
	products, err := p.rp.GetProducts()
	if err != nil {
		return nil, err
	}
	fmt.Println("GetProducts", products)
	return products, nil
}
