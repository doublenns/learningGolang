package main

import "fmt"

func main() {
    fmt.Println("Please enter number a number to be doubled:")
    var inputNum float64
    fmt.Scanf("%f", &inputNum)

    outputNum := inputNum * 2
    fmt.Println(outputNum)
}
