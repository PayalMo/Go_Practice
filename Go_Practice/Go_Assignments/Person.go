package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastname string
	age      int
	address  string
}

func main2() {
	PersonDetails()
}

func PersonDetails() {
	p := Person{name: "Payal", lastname: "More", age: 30, address: ""}
	fmt.Println("Person Details:", p)

	fmt.Println("Full Name:", fullname(p))
	fmt.Printf("Formatted Full name: %s %s %d \n", p.name, p.lastname, p.age)

	fmt.Println("Update Last Name:", p.updateLastName())

	fmt.Println("Full Name after updation:", fullname(p))

	fmt.Printf("%s is %d years old.And %s is in %s Categaory\n", p.name, p.age, p.name, p.ageCategory(p.age))

	fmt.Println("Is Address Available:", checkIsAddressAvailable(p))

}

func fullname(p Person) string {
	return p.name + " " + p.lastname
}

func (p Person) updateLastName() string {
	p.lastname = "Patil"
	return p.name + " " + p.lastname
}

func (p Person) ageCategory(age int) string {
	switch {
	case age < 0:
		return "Invalid Age"
	case age < 13:
		return "Child"
	case age < 20:
		return "Teenager"
	case age < 60:
		return "Adult"
	default:
		return "Senior Citizen"
	}
}

func checkIsAddressAvailable(p Person) string {
	if p.address == "" {
		return "Address Not Available"
	}
	return "Address Available"
}
