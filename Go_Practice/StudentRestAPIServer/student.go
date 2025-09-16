package main

type Student struct {
	RollNo    int    `json:"roll_no"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address   string `json:"address"`
	Age       int    `json:"age"`
}

// In-memory data store
var students []Student
