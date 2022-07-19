package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func openTextFile(fileName string) string {
	value, _ := ioutil.ReadFile(fileName)
	text := string(value)

	return text
}

func openJsonFile(fileName string) (*os.File, error) {
	return os.Open(fileName)
}

func unmarshalJson(file *os.File, v any) {
	byteValue, _ := ioutil.ReadAll(file)

	json.Unmarshal(byteValue, &v)
}

func writeSettingsJsonFile(settings AppSettings) {
	pretty, _ := json.MarshalIndent(settings, "", " ")
	ioutil.WriteFile(SETTINGS_FILE, pretty, 0644)
}
