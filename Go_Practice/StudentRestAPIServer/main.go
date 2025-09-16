package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	mainAPIData()
}

func mainAPIData() {
	// Initialize router
	r := mux.NewRouter()

	// Sample data
	students = append(students,
		Student{RollNo: 1, FirstName: "Payal", LastName: "More", Address: "Pune", Age: 10},
		Student{RollNo: 2, FirstName: "Punam", LastName: "Kore", Address: "Mumbai", Age: 20})

	// Routes
	r.HandleFunc("/students", GetStudents).Methods("GET")
	r.HandleFunc("/students/{roll_no}", GetStudent).Methods("GET")
	r.HandleFunc("/students", CreateStudent).Methods("POST")
	r.HandleFunc("/students/{roll_no}", UpdateStudent).Methods("PUT")
	r.HandleFunc("/students/{roll_no}", PatchStudent).Methods("PATCH")
	r.HandleFunc("/students/{roll_no}", DeleteStudent).Methods("Delete")

	// Start server
	log.Println("Server running on http://localhost:2000")
	log.Fatal(http.ListenAndServe(":2000", r))
}
