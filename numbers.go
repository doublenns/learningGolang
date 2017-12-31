package main

import "fmt"

func main() {
    var inputNums [2]float64

    fmt.Println("Please enter 2 numbers to both be doubled.")
    fmt.Println("Number 1:")
    fmt.Scanf("%f", &inputNums[0])
    fmt.Println("Number2:")
    fmt.Scanf("%f", &inputNums[1])

    fmt.Println("Numbers doubled:")
    for _, num := range inputNums  {
        fmt.Println(num * 2)
    }

    //for i in (outputNum1, outputNum2) {
    //    fmt.Println(i)
    //}
}
