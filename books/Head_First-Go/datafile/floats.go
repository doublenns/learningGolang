// Package datafile allows reading data samples from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64 // Declare the slice to be returned
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64) // Convert file line to float64
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	// If we got this far, there were no errors, so return the array of
	// numbers and a "nil" error
	return numbers, nil
}
