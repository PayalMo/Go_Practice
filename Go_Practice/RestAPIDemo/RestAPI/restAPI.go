package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func mainAPIData() {
	// Initialize router
	r := mux.NewRouter()

	// Sample data
	employees = append(employees,
		Employeee{ID: "1", FirstName: "Raju", LastName: "Rastogi", Email: "raju@gmail.com"},
		Employeee{ID: "2", FirstName: "Jenny", LastName: "M", Email: "jenny@gmail.com"})

	// Routes
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
	r.HandleFunc("/employees", CreateEmployee).Methods("POST")

	// Start server
	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
