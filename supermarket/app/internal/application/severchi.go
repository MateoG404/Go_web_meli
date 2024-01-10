// This script contains the chi router for the server application and their configuration

package application

import (
	"fmt"
	"net/http"
	internal "supermarket/app/internal"
	handlers "supermarket/app/internal/handlers"

	services "supermarket/app/internal/services"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/mux"
)

// Struct to create the server
type Server struct {
	Port   string
	Router *mux.Router

	Address string
}

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// Function to create the server in the localhost
func (s *Server) CreateServer(addres string) *Server {
	address := "localhost"
	port := "8080"

	// Check if the address and port are empty or not and assign them
	if s.Address != "" {
		address = s.Address
	}
	if s.Port != "" {
		port = s.Port
	}

	// Create the server
	return &Server{
		Router:  mux.NewRouter(),
		Address: address,
		Port:    port,
	}
}
func (s *Server) Run() error {
	fmt.Println(green("Running the server..."))

	// Configurate Dependencies
	// Configurate Repositories
	rp := internal.NewProductsRepository()

	// Configurate Services
	sv := services.NewProductsDefaultService(*rp)

	// Configurate Handlers
	hd := handlers.NewProductsDefault(sv)

	// Create the router
	router := chi.NewRouter()

	// Use the ping handler
	router.Get("/ping", handlers.PingHandler)

	// Use the products handler
	router.Get("/products", hd.ProductsHandler)

	// Run the server and listen to the port
	http.ListenAndServe(s.Address+":"+s.Port, s.Router)
	return nil

}
