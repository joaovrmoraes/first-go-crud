package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaovrmoraes/first-go-crud/database"
	"github.com/joaovrmoraes/first-go-crud/handlers"
)

func main() {
	router := mux.NewRouter()

	fmt.Println("Running in port 8080 🚀")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateUser(db, w, r)
	}).Methods("POST")

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListUsers(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.ListUserById(db, w, r)
	}).Methods("GET")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateUser(db, w, r)
	}).Methods("PUT")

	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteUser(db, w, r)
	}).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
