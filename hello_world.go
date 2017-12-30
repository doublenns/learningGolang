package main

import "fmt"
import "os"

// This is a single line Go comment

/*
This is a
multi-line Go comment
*/

func main() {
    x := "Whaddup "
    x += "doe!!"
    fmt.Println(x)
    fmt.Println("Hello World"[1])
    os.Exit(0)
}
