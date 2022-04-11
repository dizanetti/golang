package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var allPerson []person

func homeLink(w http.ResponseWriter, r *http.Request) {
	log.Println("Root Path")

	fmt.Fprintf(w, "Welcome to my first API-REST with golang!")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Create a new Person")

	var newPerson person

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &newPerson)

	allPerson = append(allPerson, newPerson)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newPerson)
}

func getOnePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Find a Person")

	personID := mux.Vars(r)["id"]

	for _, singlePerson := range allPerson {
		if singlePerson.Identification == personID {
			json.NewEncoder(w).Encode(singlePerson)
		}
	}
}

func getAllPersons(w http.ResponseWriter, r *http.Request) {
	log.Println("List all Person")

	json.NewEncoder(w).Encode(allPerson)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Update a Person data")

	personID := mux.Vars(r)["id"]
	var updatedPerson person

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error")
	}

	json.Unmarshal(reqBody, &updatedPerson)

	for i, singlePerson := range allPerson {
		if singlePerson.Identification == personID {
			singlePerson.Title = updatedPerson.Title
			singlePerson.Description = updatedPerson.Description
			allPerson = append(allPerson[:i], singlePerson)
			json.NewEncoder(w).Encode(singlePerson)
		}
	}
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete a Person")

	personID := mux.Vars(r)["id"]

	for i, singlePerson := range allPerson {
		if singlePerson.Identification == personID {
			allPerson = append(allPerson[:i], allPerson[i+1:]...)
			fmt.Fprintf(w, "The Person with ID %v has been deleted successfully", personID)
		}
	}
}
