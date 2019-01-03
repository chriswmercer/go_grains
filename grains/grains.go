//Package grains provides a function to get the number of grains on a square
package grains

import (
	"errors"
)

//Square returns the number of grains for a given square
func Square(input int) (uint64, error) {
	if input > 64 || input < 1 {
		return 0, errors.New("input value must be between 1 and 64 inclusive")
	}

	return 1 << (uint(input - 1)), nil
}

//Total returns the total number of grains
func Total() uint64 {
	return 1<<64 - 1
}
