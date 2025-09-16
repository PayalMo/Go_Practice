package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	initDB()

	r := mux.NewRouter()
	r.HandleFunc("/students", getStudents).Methods("GET")
	r.HandleFunc("/students/{roll_no}", getStudent).Methods("GET")
	r.HandleFunc("/students", createStudent).Methods("POST")
	r.HandleFunc("/students/{roll_no}", updateStudent).Methods("PUT")
	r.HandleFunc("/students/{roll_no}", patchStudent).Methods("PATCH")
	r.HandleFunc("/students/{roll_no}", deleteStudent).Methods("DELETE")

	fmt.Println("🚀 Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println()

}

// connect to DB
func initDB() {
	connStr := "host=127.0.0.1 port=5432 user=postgres password=yourpassword dbname=studentdb sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("sql.Open: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	fmt.Println("✅ Connected to PostgreSQL")
}
