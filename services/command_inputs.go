package services

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parkinglot/models"
	"strconv"
	"strings"
)

// CommandInput is struct command input
type CommandInput struct {
	typ         models.CommandInputTypes
	fileNameOpt string
}

// ParkingLotCommandInput is struct parking lot command input
type ParkingLotCommandInput struct {
	reader io.Reader

	parkingLotSvc IParking
}

// NewCommandInput is new command input instance
func NewCommandInput() *CommandInput {
	return &CommandInput{}
}

func newParkingLotCommandInput(reader io.Reader, parkingLotSvc IParking) *ParkingLotCommandInput {
	return &ParkingLotCommandInput{
		reader:        reader,
		parkingLotSvc: parkingLotSvc,
	}
}

// printf wrap print formatter
func printf(topic string, params ...interface{}) {
	fmt.Printf(fmt.Sprintf("%s\n", topic), params...)
}

// Type is command input type
func (svc *CommandInput) Type(typ models.CommandInputTypes) *CommandInput {
	svc.typ = typ
	return svc
}

// FileName is set for type file input type
func (svc *CommandInput) FileName(name string) *CommandInput {
	svc.fileNameOpt = name
	return svc
}

// Run is command input type
func (svc *CommandInput) Run() {
	var reader io.Reader
	switch svc.typ {
	case models.FileType:
		file, err := os.Open(svc.fileNameOpt)
		if err != nil {
			printf("File is invalid (%s)", err.Error())
			os.Exit(1)
		}

		reader = bufio.NewReader(file)
	default:
		reader = bufio.NewReader(os.Stdin)
	}

	parkingLotSvc := NewParking("parking-lot")
	parkingCommand := newParkingLotCommandInput(reader, parkingLotSvc)
	parkingCommand.start()
}

func (svc *ParkingLotCommandInput) start() {
	if svc.reader == nil || svc.parkingLotSvc == nil {
		printf("Internal server error")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(svc.reader)
	for scanner.Scan() {
		cmdStrs := strings.TrimSpace(scanner.Text())
		svc.commands(cmdStrs)
	}
}

func (svc *ParkingLotCommandInput) commands(cmdStrs string) {
	cmds := strings.Split(cmdStrs, " ")
	if len(cmds) == 0 {
		return
	}
	command := models.ParkingLotCommandInputs(cmds[0])
	attributes := cmds[1:]

	switch command {
	case models.CreateParkingLot:
		svc.handleCreateParkingLot(attributes...)
	case models.ParkInLot:
		svc.handleParkInLot(attributes...)
	case models.LeaveFromLot:
		svc.handleLeaveFromLot(attributes...)
	case models.GetBusyParkingStatus:
		svc.handleGetBusyParkingStatus(attributes...)
	case models.GetPlateNoByCarColor:
		svc.handleGetPlateNoByCarColor(attributes...)
	case models.GetLotNoByCarColor:
		svc.handleGetLotNoByCarColor(attributes...)
	case models.GetLotNoByPlateNo:
		svc.handleGetLotNoByPlateNo(attributes...)
	case models.Exit:
		os.Exit(0)
	default:
		printf("Invalid parking lot command.")
	}
}

func (svc *ParkingLotCommandInput) handleCreateParkingLot(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc
	if len(attrs) == 0 {
		printf("Please input parking lot amount")
		return
	}

	parkingLotAmountStr := attrs[0]
	parkingLotAmount, err := strconv.Atoi(parkingLotAmountStr)
	if err != nil {
		printf("Please input parking lot number after parking lot command")
		return
	}
	isCreated, err := parkingLotSvc.CreateParkingLot(parkingLotAmount)
	if err != nil {
		printf(err.Error())
		return
	}

	if !isCreated {
		printf("Cannot create parking lot")
		return
	}

	printf("Created a parking lot with %d slots", parkingLotAmount)
}

func (svc *ParkingLotCommandInput) handleParkInLot(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc
	if len(attrs) == 0 {
		printf("Please input registration number and car color (optional)")
		return
	}

	plateNumber := attrs[0]
	color := ""
	if len(attrs) == 2 {
		color = attrs[1]
	}

	car := NewCar(plateNumber, color, 1)
	parkLot, err := parkingLotSvc.Park(car)
	if err != nil {
		printf(err.Error())
		return
	}
	if parkLot == nil {
		printf("Cannot park at parking lot")
		return
	}

	printf("Allocated slot number: %d", parkLot.lotNo)
}

func (svc *ParkingLotCommandInput) handleLeaveFromLot(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc
	if len(attrs) == 0 {
		printf("Please input registration number and car color (optional)")
		return
	}

	lotNoStr := attrs[0]
	lotNo, err := strconv.Atoi(lotNoStr)
	if err != nil {
		printf("Please input parking lot number after a command")
		return
	}

	isLeave, err := parkingLotSvc.Leave(lotNo)
	if err != nil {
		printf(err.Error())
		return
	}
	if !isLeave {
		printf("Cannot leave at parking lot")
		return
	}

	printf("Slot number %d is free", lotNo)
}

func (svc *ParkingLotCommandInput) handleGetBusyParkingStatus(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc

	parkingLotSvc.BusyStatusTable()
}

func (svc *ParkingLotCommandInput) handleGetPlateNoByCarColor(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc

	if len(attrs) == 0 {
		printf("Please input car colour")
		return
	}

	parkingLots := parkingLotSvc.GetParkingLotsWithCarColor(attrs[0])

	plateNos := []string{}
	for _, parkingLot := range parkingLots {
		if parkingLot.vehicle == nil {
			continue
		}
		plateNos = append(plateNos, parkingLot.vehicle.PlateNumber())
	}

	printf("%s", strings.Join(plateNos, ", "))
}

func (svc *ParkingLotCommandInput) handleGetLotNoByCarColor(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc

	if len(attrs) == 0 {
		printf("Please input car colour")
		return
	}

	parkingLots := parkingLotSvc.GetParkingLotsWithCarColor(attrs[0])

	lotNos := []string{}
	for _, parkingLot := range parkingLots {
		lotNos = append(lotNos, fmt.Sprintf("%d", parkingLot.lotNo))
	}

	printf("%s", strings.Join(lotNos, ", "))
}

func (svc *ParkingLotCommandInput) handleGetLotNoByPlateNo(attrs ...string) {
	parkingLotSvc := svc.parkingLotSvc

	if len(attrs) == 0 {
		printf("Please input registration number")
		return
	}

	parkingLots := parkingLotSvc.GetParkingLotsWithPlateNo(attrs[0])

	plateNos := []string{}
	for _, parkingLot := range parkingLots {
		plateNos = append(plateNos, fmt.Sprintf("%d", parkingLot.lotNo))
	}

	if len(plateNos) == 0 {
		printf("Not found")
		return
	}

	printf("%s", strings.Join(plateNos, ", "))
}
