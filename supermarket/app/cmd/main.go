package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"supermarket/app/internal"
	"supermarket/app/internal/application"
	handler "supermarket/app/internal/handlers"
	"supermarket/app/internal/services"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// Function created to load the JSON file into the program and return a slice of strings
// Function created to load the JSON file into the program and return a slice of Products
func loadJSON(path string) ([]internal.Products, error) {
	// slice to save the data
	var p []internal.Products

	// Read the file
	file, err := ioutil.ReadFile(path)

	// Check if the file was read correctly or the file exists
	if err != nil {
		return nil, err
	}

	// Unmarshal the data into the slice
	err = json.Unmarshal(file, &p)
	if err != nil {
		return nil, err
	}

	// Return the slice of Products
	return p, nil
}
func main() {

	// _________________________________

	// Solution for the first task
	fmt.Println(green("Leyendo los archivos JSON..."))
	data, err := loadJSON("/Users/agutierrezme/Desktop/MateoCodes/Go_web_meli/supermarket/app/internal/data/products.json")

	if err != nil {
		fmt.Println(red(err))

	}
	// Save all the data in the ProductsRepository
	// Save all the data in the ProductsRepository
	productsRepository := internal.NewProductsRepository()

	for _, product := range data {
		productsRepository.AddNewProduct(product)
	}
	fmt.Println(green("Los archivos JSON se leyeron correctamente"))

	// _________________________________

	// Create the products service with the repository
	productsService := services.NewProductsDefaultService(*productsRepository)
	productsHandler := handler.NewProductsDefault(productsService)

	// Solution for the second task

	// Create the server
	server := &application.Server{}
	server = server.CreateServer("")
	server.Router.HandleFunc("/products", productsHandler.ProductsHandler).Methods("GET")
	server.Router.HandleFunc("/products/id", productsHandler.ProductByIDHandler).Methods("GET")
	server.Run()

}
