package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Simple shell")
	fmt.Println("---------------------")
	for {
		fmt.Print("Enter any text here: ")
		text, _ := reader.ReadString('\n')
		fmt.Print("You entered", text)

		// convert CRLF to LF
		// text = strings.Replace(text, LineBreak, "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
	}
}
