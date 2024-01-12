// This script contains the services for the products handler
package services

import (
	"fmt"
	"supermarket/app/internal"
	repository "supermarket/app/internal/repository"
	"time"
)

var (
	ErrInvalidDateFormat = fmt.Errorf("invalid date format")
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

// AddNewProductInput is the service to add a new product using the client input

func (p *ProductsDefault) AddNewProductInput(product internal.Products) (err error) {
	// External services
	// ...
	// -add the new product to the repository
	p.rp.AddNewProduct(product)
	return
}

// Function to validate the business logic for the post method

func (p *ProductsDefault) ValidateProductBussinessLogic(product internal.Products) (err error) {

	// Verify if the ID is already in the repository (if it is, return an error)
	new_id, existId, err := p.rp.IdExists(product.Id, product.CodeValue) // If it is true the ID already exists or is not valid so we have to end the process and return an error
	if existId {
		return err
	}
	// Set the new ID for the product
	product.Id = new_id

	// Verify if the code_value is unique (if not, return an error)
	if p.rp.VerifyCodeValue(product.CodeValue) {
		return repository.ErrProductExists
	}

	// Verify if the expiration date is valid (if not, return an error) have to be in the format xx/xx/xxxx
	_, err = time.Parse("02/01/2006", product.Expiration)
	if err != nil {
		return ErrInvalidDateFormat
	}
	return
}
