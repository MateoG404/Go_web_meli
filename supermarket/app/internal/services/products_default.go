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
