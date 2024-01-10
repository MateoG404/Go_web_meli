package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"supermarket/app/internal"
)

// Function created to load the JSON file into the program and return an slice
func loadJSON(path string) ([]string, error) {
	// slice to save the data
	var p []internal.Products

	// Read the file
	file, err := ioutil.ReadFile(path)

	// Check if the file was read correctly or the file exists
	if err != nil {
		return nil, err
	}

	// Save the data into the slice and return it
	// Unmarshal the data into the slice

	err = json.Unmarshal(file, &p)
	if err != nil {
		return nil, err
	}

	fmt.Println(p)
	return nil, nil
}
func main() {
	data, err := loadJSON("/Users/agutierrezme/Desktop/MateoCodes/Go_web_meli/supermarket/app/internal/data/products.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
