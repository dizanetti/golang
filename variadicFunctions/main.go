package main

import "fmt"

func sumValues(values ...int) {
	fmt.Print(values, " ")

	total := 0

	for _, num := range values {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sumValues(2, 3)
	sumValues(2, 10, 5, 1)
}
