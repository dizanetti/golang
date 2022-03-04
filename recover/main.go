package main

import "log"

func main() {
	defer deferFunc()

	log.Println("This func divide two value!")

	divide(0, 10)
}

func deferFunc() {
	log.Println("This is a deferred function")

	if r := recover(); r != nil {
		log.Println("Recovering from panic, error is: ", r)
	}
}

func divide(x int, y int) {
	log.Println(y / x)
}
