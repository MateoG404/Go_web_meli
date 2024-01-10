package main

import (
	"app_ejercicio1/app_ejercicio1/internal"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Function created to load the JSON file into the program and return an slice
func loadJSON(path string) ([]string, error) {
	var data internal.Products

	// Read the file
	file, err := ioutil.ReadFile(path)

	// Check if the file was read correctly or the file exists
	if err != nil {
		return nil, err
	}

	// Save the data into the slice and return it
	// Unmarshal the data into the slice

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	fmt.Println(data)
	return data, nil
}
func main() {
	data, err := loadJSON("/Users/agutierrezme/Desktop/MateoCodes/ejercicios_w11_meli/Go/03-Web/Clase-01-API/code/ejercicio_3/Go_web_meli/supermarket/app/internal/data/products.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
