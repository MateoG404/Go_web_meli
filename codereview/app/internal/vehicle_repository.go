package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)
	// FindVehicle is a method that returns a vehicle by id
	FindVehicle(id int) (v Vehicle, err error)
	// FindVehicleByRegistration is a method that returns a vehicle by registration
	FindVehicleByRegistration(registration string) (v Vehicle, err error)
	// CreateVehicle is a method that creates a vehicle
	CreateVehicle(v Vehicle) (err error)
}
