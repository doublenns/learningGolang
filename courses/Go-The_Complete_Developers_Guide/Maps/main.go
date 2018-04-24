package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// Different ways to make a Map
	// var colors map[string]string
	// colors := make(map[string]string)

	// How to add and delete items to a Map after declaration
	// colors["white"] = "#ffffff"
	// colors["red"] = "#ff0000"
	// delete(colors, "white")

	printMap(colors)
}

// How to iterate over Maps
func printMap(c map[string]string) {
	for color, hexCode := range c {
		fmt.Println("Hex code for", color, "is", hexCode)
	}
}
