package main

import "fmt"

func main() {

    i := 1
    for i <= 10 {
        fmt.Println(i)
        i += 1
    }

    for num := 1; num <= 10; num++ {
        fmt.Println(num)
    }

}
