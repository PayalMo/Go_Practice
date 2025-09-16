package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init function called")
}

func main() {
	//fmt.Println("Hello World")

	// testVariables()
	// testForLoop()
	// testSlice()
	// testArray()
	// forLoopWithIteration()

	//main1()
	//main2()
	//main3()
	//main4()
	//main5()
	//main6()
	//main7()
	//main8()
	//main9()
	//evenOddChannel()
	//main10()
	//main11()
	numAndLetters()

}

func testVariables() {

	var name string = "John"
	var age int = 30
	var isStudent bool = false

	country := "USA"
	height := 5.9

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Country:", country)
	fmt.Println("Height:", height)

	name = "Alice"
	fmt.Println("Updated Name:", name)

	const x = 10
	println("Constant x:", x)

	number, name, isTrue := 42, "The answer", true
	fmt.Println("Number:", number, "Name:", name, "IsTrue:", isTrue)

}

func testForLoop() {

	for i := 0; i <= 10; i++ {
		fmt.Println("For loop with increment:", i)
	}

	for i := range [10]int{} {
		fmt.Println(i + 1)
	}
}

// MARK: Slice
func testSlice() {
	colors := []string{"red", "green", "blue", "yellow", "purple", "orange"}

	fmt.Println("BeforeUpdate:", colors)
	fmt.Println("Length:", len(colors))
	fmt.Println("Capacity:", cap(colors))

	println("=========")

	colors[0] = "black"

	colors = append(colors, "pink")
	colors = append(colors, "white")
	fmt.Println("AfterUpdate:", colors)
	fmt.Println("Length:", len(colors))
	fmt.Println("Capacity:", cap(colors))
}

// MARK: Array
func testArray() {
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	numbers[0] = 10

	fmt.Println("Updated Array:", numbers)

}

// MARK: For Loop with Iteration number and index
func forLoopWithIteration() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
		fmt.Println("Value:", value)
	}

	for index := range numbers {
		fmt.Printf("Index: %d\n", index)
		fmt.Println("Index Value:", numbers[index])
	}

}
