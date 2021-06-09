package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") // Open the data file for reading
	// If there was an error opening the file, report it and exit
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) // Create new Scanner for the file
	// Loops until the end of file is reached & scanner.Scan returns false
	for scanner.Scan() { // Read a line from the file
		fmt.Println(scanner.Text())
	}
	err = file.Close() // Close the file to free resources
	// If there was an error closing the file, report it and exit
	if err != nil {
		log.Fatal()
	}
	// If there was an error scanning the file, report it and exit
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
