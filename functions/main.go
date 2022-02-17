package main

import "fmt"

func calculateBonus(valueOne float64, valueTwo float64) (float64, string) {
	var phrase string
	result := valueOne + valueTwo

	if result > 1000 {
		phrase = "Excelent!!!"
	} else {
		phrase = "Maybe next time!!!"
	}

	return result, phrase
}

func main() {

	// Function return two values
	valOne, _ := calculateBonus(250, 580)
	_, phrase := calculateBonus(250, 580)
	fmt.Println("Value:", valOne, "Result:", phrase)
}
