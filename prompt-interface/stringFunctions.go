package main

import "strings"

func splitString(value string, splitString string) []string {
	return strings.Split(value, splitString)
}

func findAndDelete(s []string, itemToDelete string) []string {
	var newItens [5]string
	index := 0

	for _, i := range s {
		if i != itemToDelete {
			newItens[index] = i
			index++
		}
	}
	return newItens[:index]
}
