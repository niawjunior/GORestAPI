package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person Struct
type Person struct {
	Firstname string `json:"firstname,omitempty"`
}

var people []Person

// GetPeople Function
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{Firstname: "pasupol"})
	people = append(people, Person{Firstname: "niaw"})
	router.HandleFunc("/people", GetPeople).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
