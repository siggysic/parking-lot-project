package main

import (
	"os"
	"parkinglot/models"
	"parkinglot/services"
)

func main() {
	// Input command type
	// 1. File type
	// 2. Standard input type
	cmd := services.NewCommandInput()
	if args := os.Args[1:]; len(args) > 0 {
		cmd.Type(models.FileType).FileName(args[0])
	} else {
		cmd.Type(models.InputType)
	}

	cmd.Run()
}
