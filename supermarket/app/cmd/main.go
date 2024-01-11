package main

import (
	"fmt"

	application "supermarket/app/internal/application"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func main() {
	// env
	// ...

	// app
	// - config
	app := application.CreateServer("")
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}

	/*


				server := &application.Server{}
				server = server.CreateServer("")
				server.Router.HandleFunc("/ping", handler.PingHandler).Methods("GET")
				server.Router.HandleFunc("/products", productsHandler.ProductsHandler).Methods("GET")
				server.Router.HandleFunc("/products/{id}", productsHandler.ProductByIDHandler).Methods("GET")
				server.Router.HandleFunc("/products/search", productsHandler.ProductRange).Methods("GET")
		server.Run()
	*/

}
