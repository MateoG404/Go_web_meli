package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindVehicle is a method that returns a vehicle by id
	FindVehicle(id int) (v Vehicle, err error)
	// ValidationInputData is a method that validates the input data of a vehicle
	ValidationInputData(id int, brand, model, registration, color string, fabricationYear, capacity int, maxspeed float64, fueltype, transmission string, weight, height, length, width float64) bool
	// CreateVehicle is a method that creates a vehicle
	CreateVehicle(v Vehicle) (err error)
}
