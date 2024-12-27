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

// printWithBorder prints a message with a decorative border
func printWithBorder(message string) {
	width := len(message) + 4
	border := strings.Repeat("=", width)
	fmt.Printf("\n%s\n %s \n%s\n", border, message, border)
}

// printCommandExecution prints command being executed
func printCommandExecution(command string) {
	fmt.Printf("\n▶ Executing: %s\n", command)
	fmt.Println(strings.Repeat("-", 40))
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
		return "🚫 Sorry, parking lot is full"
	}

	pl.slots[slot] = regNumber
	return fmt.Sprintf("✅ Successfully allocated slot number: %d for car: %s", slot, regNumber)
}

// Leave handles car leaving the parking lot
func (pl *ParkingLot) Leave(regNumber string, hours int) string {
	for slot, reg := range pl.slots {
		if reg == regNumber {
			delete(pl.slots, slot)
			charge := calculateCharge(hours)
			return fmt.Sprintf("🚗 Car with registration number %s left from slot %d\n💰 Parking charge: $%d for %d hours",
				regNumber, slot, charge, hours)
		}
	}
	return fmt.Sprintf("❌ Registration number %s not found in the parking lot", regNumber)
}

// Status prints current parking lot status
func (pl *ParkingLot) Status() {
	printWithBorder("Current Parking Status")

	if len(pl.slots) == 0 {
		fmt.Println("🅿️  Parking lot is empty")
		return
	}

	fmt.Printf("\n%-10s | %-15s | %-10s\n", "Slot No.", "Registration", "Status")
	fmt.Println(strings.Repeat("-", 40))

	for i := 1; i <= pl.capacity; i++ {
		if reg, exists := pl.slots[i]; exists {
			fmt.Printf("%-10d | %-15s | %-10s\n", i, reg, "Occupied")
		} else {
			fmt.Printf("%-10d | %-15s | %-10s\n", i, "-", "Available")
		}
	}

	fmt.Printf("\nTotal Capacity: %d | Occupied: %d | Available: %d\n",
		pl.capacity, len(pl.slots), pl.capacity-len(pl.slots))
}

// calculateCharge calculates parking charge based on hours
func calculateCharge(hours int) int {
	if hours <= 2 {
		return 10
	}
	return 10 + ((hours - 2) * 10)
}

func displayMenu() {
	printWithBorder("Available Commands")
	fmt.Println("1. park <registration_number>")
	fmt.Println("2. leave <registration_number> <hours>")
	fmt.Println("3. status")
	fmt.Println("4. exit")
	fmt.Println("\nEnter your choice (1-4):")
}

func runInteractiveCLI() {
	printWithBorder("Welcome to Parking Lot System")
	fmt.Print("Enter parking lot capacity: ")

	reader := bufio.NewReader(os.Stdin)
	capacityStr, _ := reader.ReadString('\n')
	capacityStr = strings.TrimSpace(capacityStr)

	capacity, err := strconv.Atoi(capacityStr)
	if err != nil || capacity <= 0 {
		fmt.Println("❌ Invalid capacity. Please enter a positive number.")
		return
	}

	parkingLot := NewParkingLot(capacity)
	fmt.Printf("🎉 Created a parking lot with %d slots\n", capacity)

	for {
		fmt.Println("\n" + strings.Repeat("-", 40))
		displayMenu()

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter car registration number: ")
			reg, _ := reader.ReadString('\n')
			reg = strings.TrimSpace(reg)
			result := parkingLot.Park(reg)
			fmt.Println(result)

		case "2":
			fmt.Print("Enter car registration number: ")
			reg, _ := reader.ReadString('\n')
			reg = strings.TrimSpace(reg)

			fmt.Print("Enter parking duration (hours): ")
			hoursStr, _ := reader.ReadString('\n')
			hoursStr = strings.TrimSpace(hoursStr)

			hours, err := strconv.Atoi(hoursStr)
			if err != nil || hours < 0 {
				fmt.Println("❌ Invalid hours. Please enter a non-negative number.")
				continue
			}

			result := parkingLot.Leave(reg, hours)
			fmt.Println(result)

		case "3":
			parkingLot.Status()

		case "4":
			printWithBorder("Thank you for using Parking Lot System")
			return

		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		// No input file provided, run interactive CLI
		runInteractiveCLI()
	} else if len(os.Args) == 2 {
		// Input file provided, run in file mode
		file, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Printf("❌ Error opening file: %v\n", err)
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

			printCommandExecution(scanner.Text())

			switch command[0] {
			case "create_parking_lot":
				capacity, _ := strconv.Atoi(command[1])
				parkingLot = NewParkingLot(capacity)
				fmt.Printf("🎉 Created a parking lot with %d slots\n", capacity)

			case "park":
				if parkingLot == nil {
					fmt.Println("❌ Parking lot not initialized")
					continue
				}
				fmt.Println(parkingLot.Park(command[1]))

			case "leave":
				if parkingLot == nil {
					fmt.Println("❌ Parking lot not initialized")
					continue
				}
				hours, _ := strconv.Atoi(command[2])
				fmt.Println(parkingLot.Leave(command[1], hours))

			case "status":
				if parkingLot == nil {
					fmt.Println("❌ Parking lot not initialized")
					continue
				}
				parkingLot.Status()
			}
		}
	} else {
		fmt.Println("❌ Invalid number of arguments")
		fmt.Println("Usage: ./parking_lot [input_file]")
	}
}
