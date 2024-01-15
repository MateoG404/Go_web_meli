// This script contains the chi router for the server application and their configuration

package application

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	internal "supermarket/app/internal"
	handlers "supermarket/app/internal/handlers"

	repository "supermarket/app/internal/repository"
	services "supermarket/app/internal/services"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
)

// Struct to create the server
type Server struct {
	address string
}

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// Function to create the server in the localhost
func CreateServer(addres string) *Server {
	// default config / values
	// ...
	if addres == "" {
		addres = ":8080"
	}

	return &Server{
		address: addres,
	}

}

// Function created to load the JSON file into the program and return a slice of strings
func loadJSON(path string) ([]internal.Products, error) {
	// slice to store the data
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

func (s *Server) Run() error {
	fmt.Println(green("Running the server..."))
	// _________________________________
	// Solution for the first task
	fmt.Println(green("Leyendo los archivos JSON..."))
	data, err := loadJSON("/Users/agutierrezme/Desktop/MateoCodes/Go_web_meli/supermarket/app/docs/db/products.json")

	if err != nil {
		fmt.Println(red("Error al leer los archivos JSON"))
		return err
	}

	// Save all the data in the ProductsRepository
	productsRepository := repository.NewProductsRepository()

	for _, product := range data {
		productsRepository.AddNewProduct(product)
	}

	fmt.Println(green("Los archivos JSON se leyeron correctamente"))

	// _________________________________

	// Configurate Dependencies

	// Configurate Services
	sv := services.NewProductsDefaultService(*productsRepository)

	// Configurate Handlers
	hd := handlers.NewProductsDefault(sv)

	// Create the router
	router := chi.NewRouter()

	// Use the ping handler
	router.Get("/ping", handlers.PingHandler)

	// Use the products handler
	router.Get("/products", hd.ProductsHandler)

	// Use the product by id handler
	router.Get("/products/{id}", hd.ProductByIDHandler)

	// Use the product by price range handler
	router.Get("/products/search", hd.ProductRange)

	// Use the new handler to add a new product to the repository
	router.Post("/products", hd.CreateProductInput)

	// Use the update handler to update a product in the repository
	router.Put("/products/{id}", hd.CreateOrUpdateProduct)

	// Run the server and listen to the port
	http.ListenAndServe(s.address, router)
	return nil

}
