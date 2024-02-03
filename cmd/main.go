package main

import (
	"gomod/database"
	"gomod/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.Init()

	router := mux.NewRouter()

	router.HandleFunc("/create/user", handlers.HandlerCreateUser).Methods("POST")
	router.HandleFunc("/delete/user/{id}", handlers.HandlerDeleteUser).Methods("DELETE")
	router.HandleFunc("/update/user/{id}", handlers.HandlerUpdateUser).Methods("PUT")
	router.HandleFunc("/users", handlers.HandlerGetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.HandlerGetUser).Methods("GET")

	http.ListenAndServe(":8080", router)

}
