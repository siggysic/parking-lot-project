package models

// ParkingStatus type of parking status
type ParkingStatus string

const (
	// Available can park
	Available ParkingStatus = "availble"
	// Busy can not park
	Busy ParkingStatus = "busy"
	// Reserve is availble but reserve specific car
	Reserve ParkingStatus = "reserve"
)
