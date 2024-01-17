package service

import (
	"app/internal"
	"app/internal/repository"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	// InvalidInputData is the error that is returned when the input data is invalid
	ErrInvalidInputData = errors.New("invalid input data")
)

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// FindVehicle is a method that returns a vehicle by id
func (s *VehicleDefault) FindVehicle(id int) (v internal.Vehicle, err error) {
	v, err = s.rp.FindVehicle(id)
	return
}

// FindVehicleById is a method that returns a vehicle by id
func (s *VehicleDefault) FindVehicleById(id int) bool {
	// find vehicle by id
	_, err := s.rp.FindVehicle(id)
	return !errors.Is(err, repository.ErrVehicleNotFound)
}

// FindVehicleByRegistration is a method that returns a vehicle by registration
func (s *VehicleDefault) FindVehicleByRegistration(registration string) bool {
	// find vehicle by registration
	_, err := s.rp.FindVehicleByRegistration(registration)
	return !errors.Is(err, repository.ErrVehicleNotFound)
}

// ValidationInputData is a method that validates the input data of a vehicle
func (s *VehicleDefault) ValidationInputData(id int, brand, model, registration, color string, fabricationYear, capacity int, maxspeed float64, fueltype, transmission string, weight, height, length, width float64) bool {

	//Validate vehicle input data
	if id == 0 || brand == "" || model == "" || registration == "" || color == "" || fabricationYear == 0 || capacity == 0 || maxspeed == 0 || fueltype == "" || transmission == "" || weight == 0 || height == 0 || length == 0 || width == 0 {
		return false
	}
	return true
}

// CreateVehicle is a method that creates a vehicle
func (s *VehicleDefault) CreateVehicle(v internal.Vehicle) (err error) {

	// Business logic

	// - validate vehicle input data
	validateData := s.ValidationInputData(v.Id, v.Brand, v.Model, v.Registration, v.Color, v.FabricationYear, v.Capacity, v.MaxSpeed, v.FuelType, v.Transmission, v.Weight, v.Height, v.Length, v.Width)
	if !validateData {
		return ErrInvalidInputData
	}

	// -validate vehicle id
	if s.FindVehicleById(v.Id) {
		// Vehicle exists with this id, we can't create a new vehicle with this id
		return repository.ErrVehicleAlreadyExists
	}
	// -Validate registration
	if s.FindVehicleByRegistration(v.Registration) {
		// Vehicle exists with this registration, we can't create a new vehicle with this registration
		return repository.ErrVehicleAlreadyExists
	}

	// create vehicle in repository
	err = s.rp.CreateVehicle(v)
	if err != nil {
		return err
	}
	return nil
}

// FindVehicleByColorYear is a method that returns a map of vehicles by color and year of fabrication using the service and the repository

func (s *VehicleDefault) FindVehicleByColorYear(color string, fabricationYear string) (v map[int]internal.Vehicle, err error) {
	// Bussines logic

	// - Convert the year of fabrication to int
	fabricationYearInt, err := strconv.Atoi(fabricationYear)
	if err != nil {
		return nil, ErrInvalidInputData
	}

	// - validate color and year of fabrication
	if color == "" || fabricationYearInt == 0 {
		return nil, ErrInvalidInputData
	}

	// Convert uppercase color title
	color = strings.Title(color)

	// Process using the repository
	fmt.Println("color", color, "fabricationYearInt", fabricationYearInt)

	// - find vehicles by color and year of fabrication
	vehicles, err := s.rp.FindVehicleByColorYear(color, fabricationYearInt)

	// - return vehicles or error
	if err != nil {
		return nil, err
	}

	return vehicles, nil
}
