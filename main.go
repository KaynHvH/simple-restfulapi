package main

import (
	"log"
	"net/http"
	"restfulAPI/db"
	"restfulAPI/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	defer db.DB.Close()

	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
