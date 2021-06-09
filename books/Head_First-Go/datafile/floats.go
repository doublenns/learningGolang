// Package datafile allows reading data samples from files
package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line of a file
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64 // Declare the array will be returning
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0 // This var tra ks which array index should assign to
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64) // Convert file line to float64
		if err != nil {
			return numbers, err
		}
		i++
	}
	err := file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}

	// If we got this far, there were no errors, so return the array of
	// numbers and a "nil" error
	return numbers, nil
}
