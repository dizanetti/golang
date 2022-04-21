package main

import "fmt"

func main() {
	fmt.Println("one")

	defer fmt.Println("three 0")
	defer fmt.Println("three 1")
	defer fmt.Println("three 2")

	fmt.Println("two")
}
