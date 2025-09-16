package main

import "testing"

func TestPersonalDetails(t *testing.T) {
	person := Person{name: "Payal", lastname: "More", age: 30, address: ""}

	if person.name != "Payal" {
		t.Errorf("Expected name to be 'Payal', but got '%s'", person.name)
	}

	if person.lastname != "More" {
		t.Errorf("Expected lastname to be 'More', but got '%s'", person.lastname)
	}

	if person.age != 30 {
		t.Errorf("Expected age to be 30, but got %d", person.age)
	}

	if person.address != "" {
		t.Errorf("Expected address to be empty, but got '%s'", person.address)
	}
}

func TestFullname(t *testing.T) {
	person := Person{name: "Payal", lastname: "More", age: 30, address: ""}
	expectedFullname := "Payal More"
	if fullname(person) != expectedFullname {
		t.Errorf("Expected fullname to be '%s', but got '%s'", expectedFullname, fullname(person))
	}
}

func TestUpdateLastName(t *testing.T) {
	person := Person{name: "Payal", lastname: "More", age: 30, address: ""}
	expectedFullnameAfterUpdate := "Payal Patil"
	if person.updateLastName() != expectedFullnameAfterUpdate {
		t.Errorf("Expected fullname after update to be '%s', but got '%s'", expectedFullnameAfterUpdate, person.updateLastName())
	}
}

func TestAgeCategory(t *testing.T) {
	person := Person{name: "Payal", lastname: "More", age: 30, address: ""}

	tests := []struct {
		age      int
		expected string
	}{
		{-1, "Invalid Age"},
		{5, "Child"},
		{15, "Teenager"},
		{30, "Adult"},
		{65, "Senior Citizen"},
	}

	for _, test := range tests {
		if person.ageCategory(test.age) != test.expected {
			t.Errorf("For age %d, expected category to be '%s', but got '%s'", test.age, test.expected, person.ageCategory(test.age))
		}
	}
}

func TestCheckIsAddressAvailable(t *testing.T) {
	personWithNoAddress := Person{name: "Payal", lastname: "More", age: 30, address: ""}
	personWithAddress := Person{name: "John", lastname: "Doe", age: 25, address: "123 Main St"}

	if checkIsAddressAvailable(personWithNoAddress) != "Address Not Available" {
		t.Errorf("Expected 'Address Not Available' for person with no address, but got '%s'", checkIsAddressAvailable(personWithNoAddress))
	}

	if checkIsAddressAvailable(personWithAddress) != "Address Available" {
		t.Errorf("Expected 'Address Available' for person with address, but got '%s'", checkIsAddressAvailable(personWithAddress))
	}
}
