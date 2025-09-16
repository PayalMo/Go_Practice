package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// ------------------ HANDLERS ------------------

// GET all students
func getStudents(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT roll_no, name, age FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var s Student
		if err := rows.Scan(&s.RollNo, &s.Name, &s.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}
	json.NewEncoder(w).Encode(students)
}

// GET student by roll_no
func getStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["roll_no"]
	id, _ := strconv.Atoi(idStr)

	var s Student
	err := db.QueryRow("SELECT roll_no, name, age FROM students WHERE roll_no=$1", id).Scan(&s.RollNo, &s.Name, &s.Age)
	if err == sql.ErrNoRows {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// POST (create)
func createStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := db.QueryRow("INSERT INTO students (name, age) VALUES ($1, $2) RETURNING roll_no", s.Name, s.Age).Scan(&s.RollNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
	fmt.Println("Added Student --", "RollNo:", s.RollNo, "Name:", s.Name, "Age:", s.Age)
}

// PUT (full update)
func updateStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["roll_no"]
	id, _ := strconv.Atoi(idStr)

	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE students SET name=$1, age=$2 WHERE roll_no=$3", s.Name, s.Age, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.RollNo = id
	json.NewEncoder(w).Encode(s)
	fmt.Println("Updated Student --", "RollNo:", s.RollNo, "Name:", s.Name, "Age:", s.Age)
}

// PATCH (partial update: only name or age)
func patchStudent(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["roll_no"]
	id, _ := strconv.Atoi(idStr)

	var patch map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	query := "UPDATE students SET "
	args := []interface{}{}
	i := 1

	if name, ok := patch["name"]; ok {
		query += fmt.Sprintf("name=$%d,", i)
		args = append(args, name)
		i++
	}
	if age, ok := patch["age"]; ok {
		query += fmt.Sprintf("age=$%d,", i)
		args = append(args, int(age.(float64))) // JSON numbers are float64
		i++
	}

	// remove trailing comma
	query = query[:len(query)-1]
	query += fmt.Sprintf(" WHERE roll_no=$%d", i)
	args = append(args, id)

	_, err := db.Exec(query, args...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updated Student
	_ = db.QueryRow("SELECT roll_no, name, age FROM students WHERE roll_no=$1", id).Scan(&updated.RollNo, &updated.Name, &updated.Age)
	json.NewEncoder(w).Encode(updated)
}

// DELETE (delete student by roll_no)
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	// Get roll_no from URL
	params := mux.Vars(r)
	rollNo, err := strconv.Atoi(params["roll_no"])
	if err != nil {
		http.Error(w, "Invalid roll number", http.StatusBadRequest)
		return
	}

	// Delete student from DB
	_, err = db.Exec("DELETE FROM students WHERE roll_no = $1", rollNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Student deleted successfully"))
	fmt.Println("Deleted Student --", "RollNo:", rollNo)
}
