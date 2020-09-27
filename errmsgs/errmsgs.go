package errmsgs

import (
	"errors"
)

// ParkingLotIsFullError is error parking lot is full
func ParkingLotIsFullError() error {
	return errors.New("Sorry, parking lot is full")
}

// VehicalNotParkingHereError is error vehical
func VehicalNotParkingHereError() error {
	return errors.New("Sorry, vehical not parking here")
}

// InternalServerError some internal error
func InternalServerError() error {
	return errors.New("Internal server error")
}
