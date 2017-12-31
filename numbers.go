package main

import "fmt"

func main() {
    fmt.Println(`Please enter 2 numbers to both be doubled.
    Number 1:`)
    var (
        inputNum1 float64
        inputNum2 float64
    )
    fmt.Scanf("%f", &inputNum1)
    fmt.Println("   Number2:")
    fmt.Scanf("%f", &inputNum2)

    outputNum1 := inputNum1 * 2
    outputNum2:= inputNum2 * 2
    fmt.Println(outputNum1)
    fmt.Println(outputNum2)
}
