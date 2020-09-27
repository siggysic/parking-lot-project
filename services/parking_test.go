package services

import (
	"parkinglot/errmsgs"
	"parkinglot/models"
	"testing"
)

func TestInputOutputNameMustValid(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	parkingName := parking.Name()

	if mockName != parkingName {
		t.Errorf("Parking name not match")
	}
}

func TestNameParkingAllFieldsMustBeEmpty(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	parkingName := parking.Name()

	if mockName != parkingName {
		t.Errorf("Parking name not match")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}

	if !parking.IsSortAvailableLot() {
		t.Errorf("IsSort should be true")
	}
}

func TestCreateParkingLotWithInputError(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isCreated, err := parking.CreateParkingLot(-100)

	if isCreated {
		t.Errorf("Should be cannot create parking lot")
	}

	if !(err != nil && err.Error() == errmsgs.InternalServerError().Error()) {
		t.Errorf("Error should have and should be negative value")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}
}

func TestCreateParkingLotWithCreateSuccess(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isCreated, err := parking.CreateParkingLot(100)

	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	if len(parking.GetAllAvailableLotNos()) != 100 {
		t.Errorf("Get all available lot should not be 100")
	}

	if len(parking.ParkingLot()) != 100 {
		t.Errorf("All parking lot should be 100")
	}

	parkingLot := parking.ParkingLot()
	if parkingLot == nil {
		t.Errorf("Parking lot should not empty")
	}

	for lotNo, lot := range parkingLot {
		if lotNo != lot.lotNo {
			t.Errorf("Parking lot should be equal")
		}
	}

	isCreated, err = parking.CreateParkingLot(10)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	if len(parking.GetAllAvailableLotNos()) != 110 {
		t.Errorf("Get all available lot should not be 110")
	}

	if len(parking.ParkingLot()) != 110 {
		t.Errorf("All parking lot should be 110")
	}
}

func TestGetParkingLotsWithCarColorWithEmptyData(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	parkLots := parking.GetParkingLotsWithCarColor("test")

	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithCarColorWithHasParkingLotNoVehicles(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}
	parkLots := parking.GetParkingLotsWithCarColor("test")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithCarColorWithCarInParkingLotButNoColor(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithCarColor("blue")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithCarColorWithCarInParkingLotButEmptyColor(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithCarColor("")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithCarColorWithCarInParkingLotSuccessData(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithCarColor("RED")
	if parkLots == nil || len(parkLots) == 0 {
		t.Errorf("Parking lot should not empty")
	}
	if len(parkLots) != 30 {
		t.Errorf("Parking lot should be 30")
	}
}

func TestGetParkingLotsWithPlatNoWithEmptyData(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	parkLots := parking.GetParkingLotsWithPlateNo("test")

	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithPlatNoWithHasParkingLotNoVehicles(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}
	parkLots := parking.GetParkingLotsWithPlateNo("test")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithPlatNoWithCarInParkingLotButNoColor(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithPlateNo("plate-2")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithPlatNoWithCarInParkingLotButEmptyColor(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithPlateNo("")
	if parkLots != nil || len(parkLots) > 0 {
		t.Errorf("Parking lot should be empty")
	}
}

func TestGetParkingLotsWithPlatNoWithCarInParkingLotSuccessData(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	parkLots := parking.GetParkingLotsWithPlateNo("plate-1")
	if parkLots == nil || len(parkLots) == 0 {
		t.Errorf("Parking lot should not empty")
	}
	if len(parkLots) != 30 {
		t.Errorf("Parking lot should be 30")
	}
}

func TestVenicleParkingWithInputError(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	parkLot, err := parking.Park(nil)

	if parkLot != nil {
		t.Errorf("Venicle should not be parking")
	}

	if !(err != nil && err.Error() == errmsgs.ParkingLotIsFullError().Error()) {
		t.Errorf("Error should have and should be parking lot is full")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}
}

func TestVenicleLeaveParkingWithInputError(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isLeave, err := parking.Leave(-100)

	if isLeave {
		t.Errorf("Venicle should not be parking")
	}

	if !(err != nil && err.Error() == errmsgs.VehicalNotParkingHereError().Error()) {
		t.Errorf("Error should have and should be vehical not parking here")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}
}

func TestVenicleParkingWithNoParkingLot(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0
	parking := NewParking(mockName)
	vehicle := NewCar(mockPlate, mockColor, usage)

	parkLot, err := parking.Park(vehicle)

	if parkLot != nil {
		t.Errorf("Venicle should not be parking")
	}

	if !(err != nil && err.Error() == errmsgs.ParkingLotIsFullError().Error()) {
		t.Errorf("Error should have and should be parking lot is full")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}
}

func TestVenicleLeaveParkingWithNoParkingLot(t *testing.T) {
	mockName := "unit-testing"
	parking := NewParking(mockName)

	isLeave, err := parking.Leave(1)

	if isLeave {
		t.Errorf("Venicle should not be parking")
	}

	if !(err != nil && err.Error() == errmsgs.VehicalNotParkingHereError().Error()) {
		t.Errorf("Error should have and should be vehical not parking here")
	}

	if len(parking.GetAllAvailableLotNos()) > 0 {
		t.Errorf("Get all available lot should not be empty")
	}

	if parking.GetAvailableLot() != nil {
		t.Errorf("Available lot should be empty")
	}
}

func TestVenicleParkingWithSuccess(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0

	parking := NewParking(mockName)
	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	vehicle := NewCar(mockPlate, mockColor, usage)
	parkLot, err := parking.Park(vehicle)

	if parkLot == nil {
		t.Errorf("Venicle should be parking")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	if len(parking.GetAllAvailableLotNos()) != 99 {
		t.Errorf("Get all available lot should not be 99")
	}

	parkingLot := parking.ParkingLot()[1]
	if parkingLot.lotNo != 1 {
		t.Errorf("Parking lot no should be 1")
	}
	if parkingLot.status != models.Busy {
		t.Errorf("Parking lot status should be busy")
	}

	lotNots := parking.GetAllAvailableLotNos()
	if len(lotNots) != 99 {
		t.Errorf("Available lot no should be 99")
	}
	for _, lotNo := range lotNots {
		if lotNo == 1 {
			t.Errorf("Lot no 1 should not be available")
		}
	}
}

func TestVenicleParkingAndLeaveParkingWithSuccess(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0

	parking := NewParking(mockName)
	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	vehicle := NewCar(mockPlate, mockColor, usage)
	parkLot, err := parking.Park(vehicle)

	if parkLot == nil {
		t.Errorf("Venicle should be parking")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	if len(parking.GetAllAvailableLotNos()) != 99 {
		t.Errorf("Get all available lot should not be 99")
	}

	parkingLot := parking.ParkingLot()[1]
	if parkingLot.lotNo != 1 {
		t.Errorf("Parking lot no should be 1")
	}
	if parkingLot.status != models.Busy {
		t.Errorf("Parking lot status should be busy")
	}

	isLeave, err := parking.Leave(1)

	if !isLeave {
		t.Errorf("Venicle should be leave the parking")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	if len(parking.GetAllAvailableLotNos()) != 100 {
		t.Errorf("All available lot should be 100")
	}

	parkingLot = parking.ParkingLot()[1]
	if parkingLot.lotNo != 1 {
		t.Errorf("Parking lot no should be 1")
	}
	if parkingLot.status != models.Available {
		t.Errorf("Parking lot status should be available")
	}
	if parkingLot.vehicle != nil {
		t.Errorf("Vehicle in parking lot should be empty")
	}
}

func TestVenicleParkingAndLeaveParkingWithFlowSuccess(t *testing.T) {
	mockName := "unit-testing"
	mockPlate := "plate-1"
	mockColor := "red"
	var usage float32 = 1.0

	parking := NewParking(mockName)
	isCreated, err := parking.CreateParkingLot(100)
	if !isCreated {
		t.Errorf("Create Should be create success")
	}

	if err != nil {
		t.Errorf("Error should be empty")
	}

	vehicle := NewCar(mockPlate, mockColor, usage)
	testParkingHelper(parking, vehicle, 30, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 70 {
		t.Errorf("All available lot should be 70")
	}
	testLeaveHelper(parking, []int{1, 2, 3, 4, 5, 10, 11, 12, 13, 14}, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 80 {
		t.Errorf("All available lot should be 80")
	}

	testParkingHelper(parking, vehicle, 1, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 79 {
		t.Errorf("All available lot should be 79")
	}
	testLeaveHelper(parking, []int{1}, t)
	if len(parking.ParkingLot()) != 100 {
		t.Errorf("Parking lot should be 100")
	}
	if len(parking.GetAllAvailableLotNos()) != 80 {
		t.Errorf("All available lot should be 80")
	}
}

func testParkingHelper(parking IParking, vehicle IVehicle, n int, t *testing.T) {
	total := len(parking.ParkingLot()) - (len(parking.ParkingLot()) - len(parking.GetAllAvailableLotNos()))
	for i := 1; i <= n; i++ {
		parkLot, err := parking.Park(vehicle)

		if parkLot == nil {
			t.Errorf("Venicle should be parking")
		}

		if err != nil {
			t.Errorf("Error should be empty")
		}

		if len(parking.GetAllAvailableLotNos()) != (total - i) {
			t.Errorf("Get all available lot should not be %d", total-i)
		}

		parkingLot := parking.ParkingLot()[i]
		if parkingLot.lotNo != i {
			t.Errorf("Parking lot no should be %d", i)
		}
		if parkingLot.status != models.Busy {
			t.Errorf("Parking lot status should be busy")
		}

		lotNots := parking.GetAllAvailableLotNos()
		if len(lotNots) != total-i {
			t.Errorf("Available lot no should be %d", total-i)
		}
		for _, lotNo := range lotNots {
			if lotNo == i {
				t.Errorf("Lot no %d should not be available", i)
			}
		}
	}
}

func testLeaveHelper(parking IParking, lotNos []int, t *testing.T) {
	total := len(parking.GetAllAvailableLotNos())
	for _, lotNo := range lotNos {
		total++
		isLeave, err := parking.Leave(lotNo)
		if !isLeave {
			t.Errorf("Venicle should be leave the parking")
		}

		if err != nil {
			t.Errorf("Error should be empty")
		}

		if len(parking.GetAllAvailableLotNos()) != total {
			t.Errorf("All available lot should be %d", len(parking.GetAllAvailableLotNos()))
		}

		parkingLot := parking.ParkingLot()[lotNo]
		if parkingLot.lotNo != lotNo {
			t.Errorf("Parking lot no should be %d", lotNo)
		}
		if parkingLot.status != models.Available {
			t.Errorf("Parking lot status should be available")
		}
		if parkingLot.vehicle != nil {
			t.Errorf("Vehicle in parking lot should be empty")
		}
	}
}
