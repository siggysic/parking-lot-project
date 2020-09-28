package services

import (
	"fmt"
	"parkinglot/errmsgs"
	"parkinglot/models"
	"sort"
	"strings"
)

// ParkingLotKeyValue is map key value ParkingLot
type ParkingLotKeyValue map[int]*ParkingLot

// CarColorParkingLotKeyValue is map key value color and ParkingLots
type CarColorParkingLotKeyValue map[int][]*ParkingLot

// ParkingLot is keeping parking lot data
type ParkingLot struct {
	lotNo   int
	lotSize float32
	status  models.ParkingStatus
	vehicle IVehicle
}

func newParkingLot(lotNo int, lotSize float32) *ParkingLot {
	return &ParkingLot{
		lotNo:   lotNo,
		lotSize: lotSize,
		status:  models.Available,
		vehicle: nil,
	}
}

// IParking is parking interface
type IParking interface {
	Name() string

	CreateParkingLot(lotAmount int) (bool, error)

	ParkingLot() ParkingLotKeyValue

	GetAvailableLot() *ParkingLot
	GetAllAvailableLotNos() []int
	GetParkingLotsWithCarColor(color string) []*ParkingLot
	GetParkingLotsWithPlateNo(plateNo string) []*ParkingLot

	Park(vehicle IVehicle) (*ParkingLot, error)
	Leave(lotNo int) (bool, error)

	IsSortAvailableLot() bool

	BusyStatusTable()
}

// Parking is a set/get parking information
type Parking struct {
	name               string
	parkingLotKeyValue ParkingLotKeyValue

	availbleLotNos []int

	isSortAvailableLot bool
}

// NewParking is a now instant parking
func NewParking(name string) IParking {
	return &Parking{
		name:               name,
		parkingLotKeyValue: map[int]*ParkingLot{},
		availbleLotNos:     []int{},
		isSortAvailableLot: true,
	}
}

// Name is get name of parking
func (svc *Parking) Name() string {
	return svc.name
}

// ParkingLot is store data parking lots
func (svc *Parking) ParkingLot() ParkingLotKeyValue {
	return svc.parkingLotKeyValue
}

// CreateParkingLot is create parking lot with amount
func (svc *Parking) CreateParkingLot(lotAmount int) (bool, error) {
	if lotAmount <= 0 {
		return false, errmsgs.InternalServerError()
	}
	lotNo := len(svc.parkingLotKeyValue) + 1
	for i := 0; i < lotAmount; i++ {
		svc.parkingLotKeyValue[lotNo] = newParkingLot(lotNo, 1)
		svc.availbleLotNos = append(svc.availbleLotNos, lotNo)
		lotNo++
	}
	return true, nil
}

// GetAvailableLot is a get available lot that nearest
func (svc *Parking) GetAvailableLot() *ParkingLot {
	availbleLotNos := svc.GetAllAvailableLotNos()
	if len(availbleLotNos) == 0 {
		return nil
	}

	if svc.IsSortAvailableLot() {
		svc.isSortAvailableLot = false
		sort.Ints(availbleLotNos)
	}
	nearestAvalLot := availbleLotNos[0]
	return svc.ParkingLot()[nearestAvalLot]
}

// GetAllAvailableLotNos is a get available lot that nearest
func (svc *Parking) GetAllAvailableLotNos() []int {
	return svc.availbleLotNos
}

// GetParkingLotsWithCarColor is a get all parking lot that color matches
func (svc *Parking) GetParkingLotsWithCarColor(color string) []*ParkingLot {
	lastLotNo := len(svc.parkingLotKeyValue)
	var parkingLots []*ParkingLot
	for i := 1; i <= lastLotNo; i++ {
		if parkingLot, ok := svc.parkingLotKeyValue[i]; ok {
			if parkingLot.vehicle == nil {
				continue
			}
			if strings.ToLower(parkingLot.vehicle.Color()) == strings.ToLower(color) {
				parkingLots = append(parkingLots, parkingLot)
			}
		}
	}
	return parkingLots
}

// GetParkingLotsWithPlateNo is a get all vehicle that color[] matches
func (svc *Parking) GetParkingLotsWithPlateNo(plateNo string) []*ParkingLot {
	lastLotNo := len(svc.parkingLotKeyValue)
	var parkingLots []*ParkingLot
	for i := 1; i <= lastLotNo; i++ {
		if parkingLot, ok := svc.parkingLotKeyValue[i]; ok {
			if parkingLot.vehicle == nil {
				continue
			}
			if parkingLot.vehicle.PlateNumber() == plateNo {
				parkingLots = append(parkingLots, parkingLot)
			}
		}
	}
	return parkingLots
}

// Park is car park at lot
func (svc *Parking) Park(vehicle IVehicle) (*ParkingLot, error) {
	parkLot := svc.GetAvailableLot()
	if parkLot == nil {
		return nil, errmsgs.ParkingLotIsFullError()
	}

	updateParkLot := svc.parkingInLot(parkLot, vehicle)
	if !updateParkLot {
		return nil, errmsgs.InternalServerError()
	}

	return parkLot, nil
}

// Leave is car leave out of lot
func (svc *Parking) Leave(lotNo int) (bool, error) {
	isLeaved := svc.leaveFromLot(lotNo)
	if !isLeaved {
		return false, errmsgs.VehicalNotParkingHereError()
	}
	return isLeaved, nil
}

// IsSortAvailableLot is optimize for sort when needed
func (svc *Parking) IsSortAvailableLot() bool {
	return svc.isSortAvailableLot
}

// parkingInLot is local function for update park lot
func (svc *Parking) parkingInLot(carPark *ParkingLot, vehicle IVehicle) bool {
	if carPark == nil || vehicle == nil {
		return false
	}

	if carPark.vehicle != nil {
		return false
	}

	// If Lot not equal availble will not update
	// Not support Reserve status yet
	if carPark.status != models.Available {
		return false
	}

	// Update available lot stores
	// Require sort available parking lots if needed
	indParkingLotNo := 0
	for ind, lotNo := range svc.availbleLotNos {
		if lotNo == carPark.lotNo {
			indParkingLotNo = ind
			if ind != 0 {
				svc.isSortAvailableLot = true
			}
			break
		}
	}
	copy(svc.availbleLotNos[indParkingLotNo:], svc.availbleLotNos[indParkingLotNo+1:])
	svc.availbleLotNos[len(svc.availbleLotNos)-1] = -1
	svc.availbleLotNos = svc.availbleLotNos[:len(svc.availbleLotNos)-1]

	// Update struct
	carPark.vehicle = vehicle
	carPark.status = models.Busy

	return true
}

func (svc *Parking) leaveFromLot(lotNo int) bool {
	parkingLot, ok := svc.parkingLotKeyValue[lotNo]
	if !ok {
		return false
	}

	if parkingLot == nil || parkingLot.vehicle == nil {
		return false
	}

	svc.isSortAvailableLot = true
	svc.availbleLotNos = append(svc.availbleLotNos, parkingLot.lotNo)

	// Update struct
	parkingLot.vehicle = nil
	parkingLot.status = models.Available
	svc.parkingLotKeyValue[lotNo] = parkingLot

	return true
}

// BusyStatusTable is format print string
func (svc *Parking) BusyStatusTable() {
	fmt.Printf("%-12s%-19s%s\n", "Slot No.", "Registration No", "Colour")
	lastLotNo := len(svc.parkingLotKeyValue)
	for i := 1; i <= lastLotNo; i++ {
		if parkingLot, ok := svc.parkingLotKeyValue[i]; ok {
			if parkingLot.status != models.Busy || parkingLot.vehicle == nil {
				continue
			}
			fmt.Printf("%-12d%-19s%s\n", parkingLot.lotNo, parkingLot.vehicle.PlateNumber(), parkingLot.vehicle.Color())
		}
	}
}
