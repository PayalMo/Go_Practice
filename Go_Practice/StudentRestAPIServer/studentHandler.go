package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// //----- -MARK: Handlers -------////
// MARK: Get StudentList(GET Method)
func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// MARK: Get StudentByRollNo(GET Method)
func GetStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	rollNoStr := params["roll_no"]
	rollNo, err := strconv.Atoi(rollNoStr)
	if err != nil {
		http.Error(w, "Invalid roll number", http.StatusBadRequest)
		return
	}
	for _, student := range students {
		if student.RollNo == rollNo {
			json.NewEncoder(w).Encode(student)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}

// MARK: Create NewStudent(POST Method)
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	_ = json.NewDecoder(r.Body).Decode(&student)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}

// MARK: Update Student (PUT Method)
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["roll_no"] // from route /students/{roll_no}
	rollNo, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Roll Number", http.StatusBadRequest)
		return
	}

	var updatedStudent Student
	if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Ensure the roll number in URL and body are consistent
	if updatedStudent.RollNo != rollNo {
		http.Error(w, "Roll number mismatch", http.StatusBadRequest)
		return
	}

	// Find and update the student
	for i, s := range students {
		if s.RollNo == rollNo {
			students[i] = updatedStudent // replace existing student
			json.NewEncoder(w).Encode(updatedStudent)
			return
		}
	}

	// If not found
	http.Error(w, "Student not found", http.StatusNotFound)
}

// MARK: Partial Update Student (PATCH Method)
func PatchStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["roll_no"] // from URL
	rollNo, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Roll Number", http.StatusBadRequest)
		return
	}

	// Partial data structure
	var patchData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&patchData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	fmt.Println("PatchData", patchData)

	// Find the student
	for i, s := range students {
		if s.RollNo == rollNo {
			// Update only the fields provided
			fmt.Println("Before update:", students[i])

			// if name, ok := patchData["first_name"].(string); ok {
			// 	students[i].FirstName = name
			// }
			
			// if lastName, ok := patchData["last_name"].(string); ok {
			// 	students[i].LastName = lastName
			// }

			if ageVal, ok := patchData["age"]; ok {
				fmt.Printf("Raw ageVal: %#v (type %T)\n", ageVal, ageVal)
				switch v := ageVal.(type) {
				case float64:
					students[i].Age = int(v)

				case string:
					if ageInt, err := strconv.Atoi(v); err == nil {
						students[i].Age = ageInt

					}
				}
			}
			fmt.Println("After update:", students[i])
			json.NewEncoder(w).Encode(students[i])
			return
		}
	}

	http.Error(w, "Student not found", http.StatusNotFound)
}

// MARK: DELETE Student (DELETE Method)
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := mux.Vars(r)["roll_no"]
	rollNo, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid RollNo", http.StatusBadRequest)
		return
	}

	for i, s := range students {
		if s.RollNo == rollNo {
			students = append(students[:i], students[i+1:]...) // remove element
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Student not found", http.StatusNotFound)
}
