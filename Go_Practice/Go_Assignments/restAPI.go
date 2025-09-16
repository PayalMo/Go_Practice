package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Employeee struct
type Employeee struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// In-memory data store
var employees []Employeee

// Handlers
func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, emp := range employees {
		if emp.ID == params["id"] {
			json.NewEncoder(w).Encode(emp)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employeee
	_ = json.NewDecoder(r.Body).Decode(&emp)
	employees = append(employees, emp)
	json.NewEncoder(w).Encode(emp)
}

func main11() {
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
