package main

import "fmt"

func main() {
	fmt.Println("one")

	defer fmt.Println("three")

	fmt.Println("two")
}
