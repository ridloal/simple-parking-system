package main

import (
	"strings"
	"testing"
)

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

func TestCalculateCharge(t *testing.T) {
	tests := []struct {
		name     string
		hours    int
		expected int
	}{
		{"Within first 2 hours", 1, 10},
		{"Exactly 2 hours", 2, 10},
		{"One additional hour", 3, 20},
		{"Three additional hours", 5, 40},
		{"Full day parking", 24, 230},
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

func TestParkingLotOperations(t *testing.T) {
	t.Run("Complete parking operation flow", func(t *testing.T) {
		pl := NewParkingLot(3)

		// Test parking
		result1 := pl.Park("B-1234-XYZ")
		if !strings.Contains(result1, "Successfully allocated slot number: 1") {
			t.Errorf("Expected successful parking, got: %s", result1)
		}

		// Test status (indirectly through slots map)
		if len(pl.slots) != 1 {
			t.Errorf("Expected 1 occupied slot, got %d", len(pl.slots))
		}

		// Test leaving
		result2 := pl.Leave("B-1234-XYZ", 3)
		if !strings.Contains(result2, "left from slot 1") {
			t.Errorf("Expected successful leave, got: %s", result2)
		}

		// Test slot is freed
		if len(pl.slots) != 0 {
			t.Errorf("Expected 0 occupied slots after leave, got %d", len(pl.slots))
		}
	})
}

func TestEdgeCases(t *testing.T) {
	t.Run("Zero capacity parking lot", func(t *testing.T) {
		pl := NewParkingLot(0)
		result := pl.Park("KA-01-HH-1234")
		expected := "🚫 Sorry, parking lot is full"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})

	t.Run("Multiple operations on same slot", func(t *testing.T) {
		pl := NewParkingLot(1)
		pl.Park("KA-01-HH-1234")
		pl.Leave("KA-01-HH-1234", 1)
		result := pl.Park("KA-01-HH-5678")
		expected := "✅ Successfully allocated slot number: 1 for car: KA-01-HH-5678"
		if result != expected {
			t.Errorf("Expected '%s', got '%s'", expected, result)
		}
	})
}
