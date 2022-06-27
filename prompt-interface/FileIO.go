package main

import (
	"io/ioutil"
)

func openTextFile(fileName string) string {
	value, _ := ioutil.ReadFile(fileName)
	text := string(value)

	return text
}
