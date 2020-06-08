package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

// Person Definition
type Person struct {
	ID string `json:"id, omitempty"`
	FirstName string `json:"first_name, omitempty"`
	LastName string `json:"last_name, omitempty"`
	Address *Address `json:"address, omitempty"`
}

type Address struct {
	City string `json:"city, omitempty"`
	State string `json:"state, omitempty"`
}

// People Aray
var people []Person

func Home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func GetPeople(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request){
	// people_count := len(people) + 1
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"]{
			people = append(people[:index], people[index + 1:]...)
			return
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main(){
	// routes
	// http.HandleFunc("/", Home )
	// Server start
	// Router
	people = append(people, Person{ID: "1", FirstName: "Gonzalo", LastName: "Munoz", Address: &Address{City: "Viña del Mar", State: "Valparaiso"}})
	people = append(people, Person{ID: "2", FirstName: "Fernanda", LastName: "Munoz"})

	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":3000", router))
	
}




