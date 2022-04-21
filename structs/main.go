package main

import "fmt"

type address struct {
	loc     string
	number  int
	country string
}

type person struct {
	name string
	age  int
	add  address
}

func newPerson(name string) person {
	p := person{name: name}

	p.age = 18
	p.add.loc = "Street blá blá"
	p.add.number = 1000
	p.add.country = "Portugal"

	return p
}

func main() {
	fmt.Println(person{"Bob", 20, address{}})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50, add: address{loc: "Street", number: 182, country: "Espanha"}}
	fmt.Println(s)
}
