package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	json, err := json.Marshal(Person{Name: "Diego", Age: 35})
	if err != nil {
		fmt.Println(err)
	}

	// Publisher an object
	publisher("person", json)
	objectValue := subscriber("person")

	log.Println(objectValue)

	// Publisher a simple string
	publisher("simpleValue", "abcde123")
	simpleValue := subscriber("simpleValue")

	log.Println(simpleValue)

	// Unknow Key
	subscriber("unknowKey")
}
