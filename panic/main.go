package main

import "log"

func main() {
	log.Println("Hello world!")

	divide(0, 10)

	log.Println("Bye Bye!")
}

func divide(x int, y int) {
	log.Println(y / x)
}
