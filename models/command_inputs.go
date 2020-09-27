package models

// ParkingLotCommandInputs is command input parking lot
type ParkingLotCommandInputs string

// CommandInputTypes is command input type
type CommandInputTypes string

const (
	// CreateParkingLot use for create parking lot
	CreateParkingLot ParkingLotCommandInputs = "create_parking_lot"
	// ParkInLot use for car that need to park
	ParkInLot ParkingLotCommandInputs = "park"
	// LeaveFromLot use for car leave from park
	LeaveFromLot ParkingLotCommandInputs = "leave"
	// GetBusyParkingStatus use for get list parking lot that have busy status
	GetBusyParkingStatus ParkingLotCommandInputs = "status"
	// GetPlateNoByCarColor use for get list plate number that parking in parking lot by car color
	GetPlateNoByCarColor ParkingLotCommandInputs = "registration_numbers_for_cars_with_colour"
	// GetLotNoByCarColor use for get list lot no that parking in parking lot by car color
	GetLotNoByCarColor ParkingLotCommandInputs = "slot_numbers_for_cars_with_colour"
	// GetLotNoByPlateNo use for get list lot no that parking in parking lot by plate no
	GetLotNoByPlateNo ParkingLotCommandInputs = "slot_number_for_registration_number"
	// Exit use for exit from program
	Exit ParkingLotCommandInputs = "exit"
)

const (
	// FileType is file import type way
	FileType CommandInputTypes = "file"
	// InputType is input command type way
	InputType CommandInputTypes = "input"
)
