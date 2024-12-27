package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout
func captureOutput(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestNewParkingLot(t *testing.T) {
	t.Run("Successfully create new parking lot", func(t *testing.T) {
		pl := NewParkingLot(6)
		if pl.capacity != 6 {
			t.Errorf("Expected capacity 6, got %d", pl.capacity)
		}
		if len(pl.slots) != 0 {
			t.Errorf("Expected empty slots, got %d slots", len(pl.slots))
		}
	})

	t.Run("Create parking lot with zero capacity", func(t *testing.T) {
		pl := NewParkingLot(0)
		if pl.capacity != 0 {
			t.Errorf("Expected capacity 0, got %d", pl.capacity)
		}
	})
}

func TestPrintWithBorder(t *testing.T) {
	output := captureOutput(func() {
		printWithBorder("Test Message")
	})

	if !strings.Contains(output, "Test Message") {
		t.Errorf("Expected output to contain 'Test Message', got %s", output)
	}
	if !strings.Contains(output, "=") {
		t.Errorf("Expected output to contain border '=', got %s", output)
	}
}

func TestPrintCommandExecution(t *testing.T) {
	output := captureOutput(func() {
		printCommandExecution("test command")
	})

	if !strings.Contains(output, "test command") {
		t.Errorf("Expected output to contain 'test command', got %s", output)
	}
	if !strings.Contains(output, "▶") {
		t.Errorf("Expected output to contain arrow symbol, got %s", output)
	}
}

func TestFindNextAvailableSlot(t *testing.T) {
	t.Run("Find slot in empty lot", func(t *testing.T) {
		pl := NewParkingLot(2)
		slot, found := pl.findNextAvailableSlot()
		if !found || slot != 1 {
			t.Errorf("Expected slot 1, found: %v, slot: %d", found, slot)
		}
	})

	t.Run("Find slot in partially filled lot", func(t *testing.T) {
		pl := NewParkingLot(2)
		pl.slots[1] = "KA-01-HH-1234"
		slot, found := pl.findNextAvailableSlot()
		if !found || slot != 2 {
			t.Errorf("Expected slot 2, found: %v, slot: %d", found, slot)
		}
	})

	t.Run("No slot in full lot", func(t *testing.T) {
		pl := NewParkingLot(1)
		pl.slots[1] = "KA-01-HH-1234"
		_, found := pl.findNextAvailableSlot()
		if found {
			t.Error("Expected no slot to be found in full lot")
		}
	})
}

func TestPark(t *testing.T) {
	t.Run("Park car in empty lot", func(t *testing.T) {
		pl := NewParkingLot(2)
		result := pl.Park("KA-01-HH-1234")
		expected := "✅ Successfully allocated slot number: 1 for car: KA-01-HH-1234"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Park car in partially filled lot", func(t *testing.T) {
		pl := NewParkingLot(2)
		pl.Park("KA-01-HH-1234")
		result := pl.Park("KA-01-HH-9999")
		expected := "✅ Successfully allocated slot number: 2 for car: KA-01-HH-9999"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Park car in full lot", func(t *testing.T) {
		pl := NewParkingLot(1)
		pl.Park("KA-01-HH-1234")
		result := pl.Park("KA-01-BB-0001")
		expected := "🚫 Sorry, parking lot is full"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})
}

func TestLeave(t *testing.T) {
	t.Run("Car leaves after parking", func(t *testing.T) {
		pl := NewParkingLot(2)
		pl.Park("KA-01-HH-1234")
		result := pl.Leave("KA-01-HH-1234", 4)
		expected := "🚗 Car with registration number KA-01-HH-1234 left from slot 1\n💰 Parking charge: $30 for 4 hours"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Try to remove non-existent car", func(t *testing.T) {
		pl := NewParkingLot(2)
		result := pl.Leave("KA-01-HH-9999", 2)
		expected := "❌ Registration number KA-01-HH-9999 not found in the parking lot"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})
}

func TestStatus(t *testing.T) {
	t.Run("Empty parking lot status", func(t *testing.T) {
		pl := NewParkingLot(2)
		output := captureOutput(func() {
			pl.Status()
		})
		if !strings.Contains(output, "Parking lot is empty") {
			t.Errorf("Expected empty parking lot message, got %s", output)
		}
	})

	t.Run("Partially filled parking lot status", func(t *testing.T) {
		pl := NewParkingLot(2)
		pl.Park("KA-01-HH-1234")
		output := captureOutput(func() {
			pl.Status()
		})
		if !strings.Contains(output, "KA-01-HH-1234") || !strings.Contains(output, "Occupied") {
			t.Errorf("Expected status to show parked car, got %s", output)
		}
	})

	t.Run("Full parking lot status", func(t *testing.T) {
		pl := NewParkingLot(2)
		pl.Park("KA-01-HH-1234")
		pl.Park("KA-01-HH-5678")
		output := captureOutput(func() {
			pl.Status()
		})
		if !strings.Contains(output, "Total Capacity: 2 | Occupied: 2") {
			t.Errorf("Expected status to show full capacity, got %s", output)
		}
	})
}

func TestCalculateCharge(t *testing.T) {
	tests := []struct {
		name     string
		hours    int
		expected int
	}{
		{"Zero hours", 0, 10},
		{"One hour", 1, 10},
		{"Two hours", 2, 10},
		{"Three hours", 3, 20},
		{"Five hours", 5, 40},
		{"Ten hours", 10, 90},
		{"Twenty four hours", 24, 230},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := calculateCharge(test.hours)
			if result != test.expected {
				t.Errorf("For %d hours, expected charge $%d, got $%d",
					test.hours, test.expected, result)
			}
		})
	}
}

func TestMain(t *testing.T) {
	t.Run("No input file provided", func(t *testing.T) {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()

		os.Args = []string{"cmd"}
		output := captureOutput(func() {
			main()
		})

		if !strings.Contains(output, "Please provide input file path") {
			t.Errorf("Expected error message for missing input file, got %s", output)
		}
	})

	t.Run("Invalid input file", func(t *testing.T) {
		oldArgs := os.Args
		defer func() { os.Args = oldArgs }()

		os.Args = []string{"cmd", "nonexistent.txt"}
		output := captureOutput(func() {
			main()
		})

		if !strings.Contains(output, "Error opening file") {
			t.Errorf("Expected error message for invalid file, got %s", output)
		}
	})
}

func TestParkingLotOperations(t *testing.T) {
	t.Run("Complete parking operation flow", func(t *testing.T) {
		pl := NewParkingLot(3)

		// Test parking multiple cars
		results := []string{
			pl.Park("B-1234-XYZ"),
			pl.Park("B-5678-ABC"),
			pl.Park("B-9012-DEF"),
		}

		for i, result := range results {
			if !strings.Contains(result, "Successfully allocated") {
				t.Errorf("Expected successful parking for car %d, got: %s", i+1, result)
			}
		}

		// Test full lot
		fullResult := pl.Park("B-3333-GHI")
		if !strings.Contains(fullResult, "Sorry, parking lot is full") {
			t.Errorf("Expected lot full message, got: %s", fullResult)
		}

		// Test leaving
		leaveResult := pl.Leave("B-1234-XYZ", 3)
		if !strings.Contains(leaveResult, "left from slot 1") {
			t.Errorf("Expected successful leave, got: %s", leaveResult)
		}

		// Test parking in freed slot
		newParkResult := pl.Park("B-4444-JKL")
		if !strings.Contains(newParkResult, "Successfully allocated slot number: 1") {
			t.Errorf("Expected parking in freed slot, got: %s", newParkResult)
		}

		// Test status output
		output := captureOutput(func() {
			pl.Status()
		})
		if !strings.Contains(output, "B-4444-JKL") {
			t.Errorf("Expected status to show newly parked car, got: %s", output)
		}
	})
}
