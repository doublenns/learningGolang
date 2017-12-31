package main

import "fmt"

func main() {

    for i:=1; i <= 5; i++ {
        switch i{
            case 1: fmt.Println("One")
            case 2: fmt.Println("Two")
            case 3: fmt.Println("Three")
            case 4: fmt.Println("Four")
            case 5: fmt.Println("Five")
        }
    }

    fmt.Println()

    for num := 1; num <= 10; num++ {
        if num % 2 == 0 {
            fmt.Println(num, "Even")
        } else {
            fmt.Println(num, "Odd")
        }
    }

}
