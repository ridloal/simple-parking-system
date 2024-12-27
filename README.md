# Parking Lot System

A command-line parking lot management system implemented in Go. The system allows for automated ticketing, parking allocation, and fee calculation.

## Features

- Create parking lot with specified capacity
- Park cars with registration numbers
- Remove cars and calculate parking fees
- View current parking lot status
- Automated slot allocation (nearest to entry)
- Time-based parking fee calculation

### Updated Feature from Branch real-time-parking-cli

- Support interactive CLI and can operate without input.txt

## Prerequisites

- Go 1.16 or higher

## Project Structure

```
parking-lot/
├── main.go
├── main_test.go
├── input.txt
├── input2.txt
└── README.md
```

## Installation and Running

1. Clone the repository:
```bash
git clone https://github.com/ridloal/simple-parking-system
cd simple-parking-system
```

2. Initialize Go module:
```bash
go mod init simple-parking-system
```

3. Build the application:
```bash
go build main.go
```

4. Run the application with input file:
```bash
main input.txt
# or use the second example
main input2.txt

# or use interactive CLI (run without argument input)
main
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

Two example files are provided:
- `input.txt` - Basic parking operations
- `input2.txt` - More complex scenarios including full lot management

Check these files for example commands and expected output.

Run without argument `input.txt` to run interactive CLI.