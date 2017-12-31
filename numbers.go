package main

import "fmt"

func main() {
    fmt.Println("Please enter 2 numbers to both be doubled.")
    fmt.Println("Number 1:")

    var inputNums [2]float64

    fmt.Scanf("%f", &inputNums[0])
    fmt.Println("Number2:")
    fmt.Scanf("%f", &inputNums[1])

    fmt.Println("Numbers doubled:")
    for i := 0; i <= 1; i++ {
        fmt.Println(inputNums[i] * 2)
    }

    //for i in (outputNum1, outputNum2) {
    //    fmt.Println(i)
    //}
}
