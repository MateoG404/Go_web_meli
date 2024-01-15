// This code contains the test to verify the correct funciontalities of the handlers

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"supermarket/app/internal"
	repository "supermarket/app/internal/repository"
	services "supermarket/app/internal/services"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

// Function to configure the Request
func NewRequest(method, url string, body io.Reader, urlParams map[string]string, urlQuery map[string]string) *http.Request {
	// Create the request

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	// Add the url params to the request using the context
	// We pass the params like the context, using a map key value
	if urlParams != nil {
		// Add all the url params to the request
		ctx := chi.NewRouteContext()
		for key, value := range urlParams {
			// Add the key and value to the request
			ctx.URLParams.Add(key, value)
		}
		// Add the context to the request
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))
	}
	// Add the token to the request
	//req.Header.Set("Token", "123456")
	return req

}

// Test for the handler ProductByIDHandler
func TestProductsGetById(t *testing.T) {
	// arrange - Configuration

	// - Configurate a global database
	t.Run("Success Get product by id", func(t *testing.T) {
		// ARRANGE
		// - Configurate database

		db := map[int]internal.Products{
			72: {
				Id:          72,
				Name:        "Wine - Chianti Classico Riserva",
				Quantity:    458,
				CodeValue:   "",
				IsPublished: false,
				Expiration:  "24/03/2021",
				Price:       635.94,
			}}
		// - Configurate the repository
		rp := repository.NewProductsRepository(db)

		// - Configurate the service
		sv := services.NewProductsDefaultService(*rp)

		// Configurate the handler
		hd := NewProductsDefault(sv)

		// ACT

		// Configurate the request
		// We use the NewRequest function to create a new request using the context of the request to add the variables
		r := NewRequest(http.MethodGet, "/products/72", nil, map[string]string{"id": "72"}, nil)
		// Configurate the response
		w := httptest.NewRecorder()

		hd.ProductByIDHandler(w, r)
		// ASSERT

		// - Expected code,body and header
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// - Verify the code
		require.Equal(t, expectedCode, w.Code)
		// - Verify the body
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		var gotProduct internal.Products
		json.Unmarshal(bodyBytes, &gotProduct)

		expectedProduct := internal.Products{
			Id:          72,
			Name:        "Wine - Chianti Classico Riserva",
			Quantity:    458,
			CodeValue:   "",
			IsPublished: false,
			Expiration:  "24/03/2021",
			Price:       635.94,
		}
		fmt.Println("gotProduct", gotProduct, "expectedProduct", expectedProduct)
		require.Equal(t, expectedProduct, gotProduct)
		// - Verify the header
		require.Equal(t, expectedHeader, w.Header())

	})

	t.Run("Fail Get product by id", func(t *testing.T) {

		// ARRANGE
		db := map[int]internal.Products{
			72: {
				Id:          72,
				Name:        "Wine - Chianti Classico Riserva",
				Quantity:    458,
				CodeValue:   "",
				IsPublished: false,
				Expiration:  "24/03/2021",
				Price:       635.94,
			}}
		// - Configurate the repository
		rp := repository.NewProductsRepository(db)
		// - Configurate the service
		sv := services.NewProductsDefaultService(*rp)
		// - Configurate the handler
		hd := NewProductsDefault(sv)

		// ACT

		// Configurate the request
		// We use the NewRequest function to create a new request using the context of the request to add the variables
		r := NewRequest(http.MethodGet, "/products/73", nil, map[string]string{"id": "73"}, nil)
		w := httptest.NewRecorder()

		hd.ProductByIDHandler(w, r)

		// ASSERT

		// - Expected code,body and header
		// - Expected code, body and header
		expectedCode := http.StatusNotFound
		expectedHeader := http.Header{"Content-Type": []string{"text/plain"}}
		expectedBody := "product not found\n"
		// - Verify the code
		require.Equal(t, expectedCode, w.Code)
		// - Verify the body
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		require.Equal(t, expectedBody, string(bodyBytes))
		// - Verify the header
		require.Equal(t, expectedHeader, w.Header())
		// - Verify the body
		bodyBytes, _ = ioutil.ReadAll(w.Body)
		require.Equal(t, expectedBody, string(bodyBytes))
	})
}

// Test for the handler CreateProductInput

func TestCreateProductInput(t *testing.T) {
	// Test for the handler CreateProductInput with success and fail
	t.Run("Success Create product", func(t *testing.T) {

		// ARRANGE
		// - Configurate database
		db := map[int]internal.Products{}

		// - Configurate the repository
		rp := repository.NewProductsRepository(db)

		// - Configurate the service
		sv := services.NewProductsDefaultService(*rp)

		// - Configurate the handler

		hd := NewProductsDefault(sv)

		// ACT

		// Configurate the request

		body := strings.NewReader(`{
			"id": 501,
			"name": "prueba",
			"quantity":400,
			"code_value":"Asdff",
			"is_published":true,
			"expiration":"07/08/2003",
			"price":900.034
		}`)

		r := NewRequest(http.MethodPost, "/products", body, nil, nil)
		w := httptest.NewRecorder()
		w.Header().Set("Content-Type", "application/json")

		hd.CreateProductInput(w, r)
		// ASSERT

		// - Expected code,body and header
		// - Expected code,body and header
		expectedCode := http.StatusOK
		expectedBody, _ := ioutil.ReadAll(body) // Convert body to string

		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}

		// - Verify the code
		require.Equal(t, expectedCode, w.Code)
		// - Verify the body
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		require.Equal(t, string(expectedBody), string(bodyBytes)) // Compare as strings
		// - Verify the header-
		require.Equal(t, expectedHeader, w.Header())
	})

	// Test created to fail
	t.Run("Fail Create product", func(t *testing.T) {

		// ARRANGE
		// - Configurate database
		db := map[int]internal.Products{}

		// - Configurate the repository
		rp := repository.NewProductsRepository(db)

		// - Configurate the service
		sv := services.NewProductsDefaultService(*rp)

		// - Configurate the handler

		hd := NewProductsDefault(sv)

		// ACT

		// Configurate the request

		body := strings.NewReader(`{}`)

		r := NewRequest(http.MethodPost, "/products", body, nil, nil)
		w := httptest.NewRecorder()
		w.Header().Set("Content-Type", "application/json")

		hd.CreateProductInput(w, r)
		// ASSERT

		// - Expected code,body and header
		// - Expected code,body and header
		expectedCode := http.StatusBadRequest
		expectedBody, _ := ioutil.ReadAll(strings.NewReader(`All fields are required`)) // Convert body to string

		expectedHeader := http.Header{"Content-Type": []string{"text/plain; charset=utf-8"}}

		// - Verify the code
		require.Equal(t, expectedCode, w.Code)
		// - Verify the body
		bodyBytes, _ := ioutil.ReadAll(w.Body)
		require.Equal(t, string(expectedBody), string(bodyBytes)) // Compare as strings
		// - Verify the header-
		require.Equal(t, expectedHeader, w.Header())
	})
}
