package main

import (
	"encoding/json"
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func GetPeople(w http.ResponseWriter, r *http.Request){

}

func GetPerson(w http.ResponseWriter, r *http.Request){

}

func CreatePerson(w http.ResponseWriter, r *http.Request){

}

func main(){
	// routes
	// http.HandleFunc("/", Home )
	// Server start
	// Router
	router := mux.NewRouter()
	// endpoints
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":3000", router))
	
}




