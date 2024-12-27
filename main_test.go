package main

import (
	"strings"
	"testing"
)

func TestNewParkingLot(t *testing.T) {
	pl := NewParkingLot(6)
	if pl.capacity != 6 {
		t.Errorf("Expected capacity 6, got %d", pl.capacity)
	}
	if len(pl.slots) != 0 {
		t.Errorf("Expected empty slots, got %d slots", len(pl.slots))
	}
}

func TestPark(t *testing.T) {
	pl := NewParkingLot(2)

	// Test successful parking
	result := pl.Park("KA-01-HH-1234")
	if result != "Allocated slot number: 1" {
		t.Errorf("Expected 'Allocated slot number: 1', got '%s'", result)
	}

	// Test second car
	result = pl.Park("KA-01-HH-9999")
	if result != "Allocated slot number: 2" {
		t.Errorf("Expected 'Allocated slot number: 2', got '%s'", result)
	}

	// Test parking when full
	result = pl.Park("KA-01-BB-0001")
	if result != "Sorry, parking lot is full" {
		t.Errorf("Expected 'Sorry, parking lot is full', got '%s'", result)
	}
}

func TestLeave(t *testing.T) {
	pl := NewParkingLot(2)
	pl.Park("KA-01-HH-1234")

	// Test successful leave
	result := pl.Leave("KA-01-HH-1234", 4)
	expected := "Registration number KA-01-HH-1234 with Slot Number 1 is free with Charge $30"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test leave with non-existent car
	result = pl.Leave("KA-01-HH-9999", 2)
	if !strings.Contains(result, "not found") {
		t.Errorf("Expected 'not found' message, got '%s'", result)
	}
}

func TestCalculateCharge(t *testing.T) {
	tests := []struct {
		hours    int
		expected int
	}{
		{1, 10}, // Within first 2 hours
		{2, 10}, // Exactly 2 hours
		{3, 20}, // One additional hour
		{5, 40}, // Three additional hours
	}

	for _, test := range tests {
		result := calculateCharge(test.hours)
		if result != test.expected {
			t.Errorf("For %d hours, expected charge $%d, got $%d",
				test.hours, test.expected, result)
		}
	}
}

func TestStatus(t *testing.T) {
	pl := NewParkingLot(2)
	pl.Park("KA-01-HH-1234")
	pl.Park("KA-01-HH-9999")

	// Since Status() prints to stdout, we're just testing that the slots are
	// properly populated
	if len(pl.slots) != 2 {
		t.Errorf("Expected 2 occupied slots, got %d", len(pl.slots))
	}

	if pl.slots[1] != "KA-01-HH-1234" {
		t.Errorf("Expected slot 1 to contain KA-01-HH-1234, got %s", pl.slots[1])
	}
}
