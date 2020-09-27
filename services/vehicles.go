package services

// IVehicle is vehicle interface
type IVehicle interface {
	PlateNumber() string
	Color() string
	UsageLot() float32
}

// carBuilder is one of vehicle
type carBuilder struct {
	plateNumber string
	color       string
	usageLot    float32
}

// NewCar is a new struct car
func NewCar(plateNumber, color string, usage float32) IVehicle {
	return &carBuilder{plateNumber: plateNumber, color: color, usageLot: usage}
}

func (svc *carBuilder) PlateNumber() string {
	return svc.plateNumber
}

func (svc *carBuilder) Color() string {
	return svc.color
}

func (svc *carBuilder) UsageLot() float32 {
	return svc.usageLot
}
