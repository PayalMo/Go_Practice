package main

import (
	"fmt"
)

func main4() {

	map1 := map[string]int{}

	map1["one"] = 11
	map1["two"] = 22
	map1["three"] = 33

	fmt.Println("Before Update Map:", map1)

	map1["one"] = 111
	map1["two"] = 222
	map1["three"] = 333

	fmt.Println("After Update Map:", map1)

	map2 := map[string]map[string]string{

		"colour":  {"Primary Colour": "Red", "Secondary Colour": "Blue"},
		"fruits":  {"Fruit1": "Apple", "Fruit2": "Banana", "Fruit3": "Orange"},
		"veggies": {"greenVeggie": "Spinach", "rootVeggie": "Carrot"},
	}
	fmt.Println("Map 2 :", map2)

	for key, value := range map2 {
		//fmt.Println("\n")

		fmt.Println("----Key:", key, "Value:----", value)
		for innerKey, innerValue := range value {
			fmt.Println("Inner Key:", innerKey, "Inner Value:", innerValue)
			//fmt.Printf("Inner Key: %s Inner Value: %s \n", innerKey, innerValue)
		}

	}
}
