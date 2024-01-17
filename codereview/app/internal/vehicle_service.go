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
	// FindVehicleById is a method that returns a vehicle by id
	FindVehicleById(id int) bool
	// FindVehicleByRegistration is a method that returns a vehicle by registration
	FindVehicleByRegistration(registration string) bool
	// FindVehicleByColorYear is a method that returns all the vehicle with that color and year
	FindVehicleByColorYear(color string, fabricationYear string) (v map[int]Vehicle, err error)
	// FindVehicleByBrandYear is a method that returns all the vehicle with that brand and year
	FindVehicleByBrandYear(brand string, startYear string, endYear string) (v map[int]Vehicle, err error)
}
