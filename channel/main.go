package main

import (
	"fmt"
	"time"
)

func say(s string, done chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
	done <- "Terminei" // escreve no channel "<-"
}
func main() {
	doneHello := make(chan string)
	doneWorld := make(chan string)

	go say("Hello", doneHello) // para bloquear e apresentar um erro de deadlock é só remover a goroutine "go"
	<-doneHello                // comentar esta linha para não bloquear a thread, assim será executado as duas goroutines e ao final executar o print

	go say("world", doneWorld)
	<-doneWorld
}
