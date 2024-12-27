# Parking Lot System

A command-line parking lot management system implemented in Go. The system allows for automated ticketing, parking allocation, and fee calculation.

## Features

- Create parking lot with specified capacity
- Park cars with registration numbers
- Remove cars and calculate parking fees
- View current parking lot status
- Automated slot allocation (nearest to entry)
- Time-based parking fee calculation

## Prerequisites

- Go 1.16 or higher

## Project Structure

```
parking-lot/
├── main.go
├── main_test.go
├── input.txt
└── README.md
```

## Installation and Running

1. Clone the repository:
```bash
git clone https://github.com/ridloal/simple-parking-system
cd simple-parking-system
```

2. Build the application:
```bash
go build main.go
```

3. Run the application with input file:
```bash
./main input.txt
```

## Running Tests

To run all tests:
```bash
go test -v
```

To run specific test:
```bash
go test -v -run TestPark
```

To check test coverage:
```bash
go test -cover
```

## Input Commands

The system accepts the following commands through the input file:

- `create_parking_lot {capacity}` - Create parking lot of size n
- `park {car_number}` - Park a car
- `leave {car_number} {hours}` - Remove(Unpark) car and calculate fee
- `status` - Print status of parking slot

## Parking Fee Structure

- First 2 hours: $10
- Every additional hour: $10/hour

## Example Usage

See `input.txt` for example commands and expected output.
