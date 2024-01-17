package handler

import (
	"app/internal"
	repository "app/internal/repository"
	"app/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}

// CreateVehicle is a method that returns a handler for the route POST /vehicles

// CreateVehicle is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) CreateVehicle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entro aca")
		// Request

		// - Get the body of the request
		var body VehicleJSON

		// - Serialize the body of the request
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid vehicle data", "error": err.Error()})
			return
		}

		// Process

		// - Create the vehicle using the service to serialize the data
		vehicle := internal.Vehicle{
			Id: body.ID,
			VehicleAttributes: internal.VehicleAttributes{
				Brand:           body.Brand,
				Model:           body.Model,
				Registration:    body.Registration,
				Color:           body.Color,
				FabricationYear: body.FabricationYear,
				Capacity:        body.Capacity,
				MaxSpeed:        body.MaxSpeed,
				FuelType:        body.FuelType,
				Transmission:    body.Transmission,
				Weight:          body.Weight,
				Dimensions: internal.Dimensions{
					Height: body.Height,
					Length: body.Length,
					Width:  body.Width,
				},
			},
		}

		fmt.Println("Vehicle: ", vehicle, "ID: ", vehicle.Id)
		// Response
		// - Send the vehicle to the service to create it

		err = h.sv.CreateVehicle(vehicle)
		if err != nil {
			switch err {
			case repository.ErrVehicleAlreadyExists:
				response.JSON(w, http.StatusConflict, map[string]interface{}{"message": "Vehicle already exists", "error": err.Error()})
			case service.ErrInvalidInputData:
				response.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid vehicle data or incomplete", "error": err.Error()})
			}
			return
		}

		response.JSON(w, http.StatusCreated, map[string]interface{}{"message": "Vehicle created successfully"})
		// chi.URLParam(r, "id")
		// chi.QueryParam(r, "id")

	}
}

// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request

		// - Get the color and year from the URL

		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		// process

		// - Get the vehicles by color and year using the service

		vehicles, err := h.sv.FindVehicleByColorYear(color, year)

		// - Check if there was an error about the query search
		if err != nil {
			response.JSON(w, http.StatusNotFound, map[string]interface{}{"message": "Not exist vehicles with the color and the year", "error": err.Error()})
			return
		}

		// response
		// if there are vehicles with the color and the year, return them
		response.JSON(w, http.StatusOK, map[string]interface{}{"message": "success", "data": vehicles})

	}
}

// GetByBrandAndyear is a method that returns a handler for the route GET /vehicles/brand/{brand}/year/{year}

func (h *VehicleDefault) GetByBrandAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request

		// - Get the brand and years from the URL
		brand := chi.URLParam(r, "brand")
		start_year := chi.URLParam(r, "start_year")
		end_year := chi.URLParam(r, "end_year")

		// Process
		vehicles, err := h.sv.FindVehicleByBrandYear(brand, start_year, end_year)

		if err != nil {
			switch err {
			case repository.ErrVehicleNotFound:
				response.JSON(w, http.StatusNotFound, map[string]interface{}{"message": "Not exist vehicles with the brand between the years given", "error": err.Error()})
				return
			case service.ErrInvalidInputData:
				response.JSON(w, http.StatusBadRequest, map[string]interface{}{"message": "Invalid vehicle input data or incomplete", "error": err.Error()})
				return
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]interface{}{"message": "Internal server error", "error": err.Error()})
			}

		}
		// Response
		response.JSON(w, http.StatusOK, map[string]interface{}{"message": "success", "data": vehicles})
	}
}
