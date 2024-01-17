package repository

import (
	"app/internal"
	"errors"
)

// Global variables
var (
	// ErrVehicleNotFound is the error that is returned when a vehicle is not found
	ErrVehicleNotFound = errors.New("vehicle not found")
	// ErrVehicleAlreadyExists is the error that is returned when a vehicle already exists
	ErrVehicleAlreadyExists = errors.New("vehicle already exists")
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindVehicle is a method that returns a vehicle by id
func (r *VehicleMap) FindVehicle(id int) (v internal.Vehicle, err error) {
	v, ok := r.db[id]
	if !ok {
		err = ErrVehicleNotFound
	}
	return
}

// FindVehicleByRegistration is a method that returns a vehicle by registration
func (r *VehicleMap) FindVehicleByRegistration(registration string) (v internal.Vehicle, err error) {
	for _, value := range r.db {
		if value.Registration == registration {
			v = value
			return
		}
	}
	err = ErrVehicleNotFound
	return
}

// CreateVehicle is a method that creates a vehicle
func (r *VehicleMap) CreateVehicle(v internal.Vehicle) (err error) {
	r.db[v.Id] = v
	return nil
}

// FindVehicleByColorYear is a method that returns all the vehicle with that color and year

func (r *VehicleMap) FindVehicleByColorYear(color string, fabricationYear int) (v map[int]internal.Vehicle, err error) {

	// Create an empty map of vehicles
	vehicle := make(map[int]internal.Vehicle)
	// Constant to know the amount of vehicles with that color and year
	var const_v int = 0

	// Iterate over the map of vehicles
	for _, value := range r.db {
		if value.Color == color && value.FabricationYear == fabricationYear {
			vehicle[const_v] = value
			const_v += 1
		}
	}
	if const_v == 0 {
		return nil, ErrVehicleNotFound
	}

	// Return the map of vehicles
	return vehicle, nil
}

// FindVehicleByBrandYear is a method that returns all the vehicle with that brand and year

func (r *VehicleMap) FindVehicleByBrandYear(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {

	// Create an empty map of vehicles
	vehicle := make(map[int]internal.Vehicle)

	// Constant to know the amount of vehicles with that brand and year
	var const_v int = 0

	// Iterate for all the databases
	for _, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			vehicle[const_v] = value
			const_v += 1
		}
	}

	// Return the map of vehicles or error if there are no vehicles with that brand and year
	if const_v == 0 {
		return nil, ErrVehicleNotFound
	}
	return vehicle, nil
}
