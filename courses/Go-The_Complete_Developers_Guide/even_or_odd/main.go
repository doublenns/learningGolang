package main

func main() {

	numbers := makeRange(1, 10)
	for _, number := range numbers {
		even_or_odd(number)
	}

}
