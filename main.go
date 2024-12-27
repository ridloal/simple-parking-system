package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ParkingLot represents the parking system
type ParkingLot struct {
	capacity int
	slots    map[int]string // maps slot number to registration number
}

// NewParkingLot creates a new parking lot with given capacity
func NewParkingLot(capacity int) *ParkingLot {
	return &ParkingLot{
		capacity: capacity,
		slots:    make(map[int]string),
	}
}

// findNextAvailableSlot finds the nearest empty slot
func (pl *ParkingLot) findNextAvailableSlot() (int, bool) {
	for i := 1; i <= pl.capacity; i++ {
		if _, exists := pl.slots[i]; !exists {
			return i, true
		}
	}
	return 0, false
}

// Park handles parking a new car
func (pl *ParkingLot) Park(regNumber string) string {
	slot, found := pl.findNextAvailableSlot()
	if !found {
		return "Sorry, parking lot is full"
	}

	pl.slots[slot] = regNumber
	return fmt.Sprintf("Allocated slot number: %d", slot)
}

// Leave handles car leaving the parking lot
func (pl *ParkingLot) Leave(regNumber string, hours int) string {
	for slot, reg := range pl.slots {
		if reg == regNumber {
			delete(pl.slots, slot)
			charge := calculateCharge(hours)
			return fmt.Sprintf("Registration number %s with Slot Number %d is free with Charge $%d",
				regNumber, slot, charge)
		}
	}
	return fmt.Sprintf("Registration number %s not found", regNumber)
}

// Status prints current parking lot status
func (pl *ParkingLot) Status() {
	fmt.Println("Slot No. Registration No.")
	for i := 1; i <= pl.capacity; i++ {
		if reg, exists := pl.slots[i]; exists {
			fmt.Printf("%d %s\n", i, reg)
		}
	}
}

// calculateCharge calculates parking charge based on hours
func calculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + ((hours - 2) * 10)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide input file path")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	var parkingLot *ParkingLot
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		if len(command) == 0 {
			continue
		}

		switch command[0] {
		case "create_parking_lot":
			capacity, _ := strconv.Atoi(command[1])
			parkingLot = NewParkingLot(capacity)
			fmt.Printf("Created a parking lot with %d slots\n", capacity)

		case "park":
			if parkingLot == nil {
				fmt.Println("Parking lot not initialized")
				continue
			}
			fmt.Println(parkingLot.Park(command[1]))

		case "leave":
			if parkingLot == nil {
				fmt.Println("Parking lot not initialized")
				continue
			}
			hours, _ := strconv.Atoi(command[2])
			fmt.Println(parkingLot.Leave(command[1], hours))

		case "status":
			if parkingLot == nil {
				fmt.Println("Parking lot not initialized")
				continue
			}
			parkingLot.Status()
		}
	}
}
