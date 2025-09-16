package main

import (
	"fmt"
)

type Student struct {
	firstName string
	lastName  string
	age       int
	ContactInfo
}

type ContactInfo struct {
	phone   string
	email   string
	zipCode string
}

func main3() {
	student := Student{firstName: "Sanchita", lastName: "More", age: 30, ContactInfo: ContactInfo{phone: "1234567890", email: "sm45@gmail.com", zipCode: "411033"}}
	fmt.Println("Student Details:", student)

	fmt.Println("Full Name:", student.Fullname())

	fmt.Println("After updating last name:")
	student.UpdateLastName("Patil")
	fmt.Println("Full Name after updation:", student.Fullname())

	updatedContactInfo := ContactInfo{phone: "9876543210", email: "gma23@gmail.com", zipCode: "411045"}
	fmt.Println("After Updated Contact Info:", updatedContactInfo)
}

func (s Student) Fullname() string {
	return s.firstName + " " + s.lastName
}

func (s *Student) UpdateLastName(newLastName string) {
	s.lastName = newLastName
}

func (s *Student) UpdateContactInfo(newPhone string, newEmail string, newZipCode string) {

	s.ContactInfo.phone = newPhone
	s.ContactInfo.email = newEmail
	s.ContactInfo.zipCode = newZipCode

	fmt.Println("Updated Contact Info:", s.ContactInfo)

}
