package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Initialize server ...")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/person", createPerson).Methods("POST")
	router.HandleFunc("/persons", getAllPersons).Methods("GET")
	router.HandleFunc("/persons/{id}", getOnePerson).Methods("GET")
	router.HandleFunc("/persons/{id}", updatePerson).Methods("PATCH")
	router.HandleFunc("/persons/{id}", deletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8085", router))
}
