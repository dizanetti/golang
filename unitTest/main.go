package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func main() {
	fmt.Println(Add(10, 2))
	fmt.Println(Subtract(8, 6))
}
