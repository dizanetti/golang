package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type person struct {
	Identification string `json:"identification"`
	Title          string `json:"title"`
	Description    string `json:"description"`
}

var allPerson []person

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my first API-REST with golang!")
}

func createPerson(w http.ResponseWriter, r *http.Request) {
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
	personID := mux.Vars(r)["id"]

	for _, singlePerson := range allPerson {
		if singlePerson.Identification == personID {
			json.NewEncoder(w).Encode(singlePerson)
		}
	}
}

func getAllPersons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(allPerson)
}

func updatePerson(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]
	var updatedPerson person

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the Person title and description only in order to update")
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
	personID := mux.Vars(r)["id"]

	for i, singlePerson := range allPerson {
		if singlePerson.Identification == personID {
			allPerson = append(allPerson[:i], allPerson[i+1:]...)
			fmt.Fprintf(w, "The Person with ID %v has been deleted successfully", personID)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/person", createPerson).Methods("POST")
	router.HandleFunc("/persons", getAllPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", getOnePerson).Methods("GET")
	router.HandleFunc("/persons/{id}", updatePerson).Methods("PATCH")
	router.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", router))
}
