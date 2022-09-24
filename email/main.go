package main

func main() {

	message := NewMessage(Compose())
	message.SetTo([]string{"zanetti.di@gmail.com"})

	Send(*message)
}
