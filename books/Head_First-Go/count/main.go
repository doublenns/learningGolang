// count tallies the number of times each line
// occurs within a file.
package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name, count := range counts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}

// func hardVersionUsingSlices() {
// 	lines, err := datafile.GetStrings("votes.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var names []string // will hold slide of candidate names
// 	var counts []int   // will hold slice w/ number of times each name occurs
// 	for _, line := range lines {
// 		matched := false
// 		for i, name := range names {
// 			if name == line {
// 				counts[i]++
// 				matched = true
// 			}
// 		}
// 		if matched == false {
// 			names = append(names, line)
// 			counts = append(counts, 1)
// 		}
// 	}
// 	for i, name := range names {
// 		fmt.Printf("%s: %d\n", name, counts[i])
// 	}
// }
