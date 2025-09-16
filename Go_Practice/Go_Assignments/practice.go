package main

import (
	"fmt"
)

func main10() {
	a := [...]int{0, 1, 2, 3}
	fmt.Println("a:", a)

	x := a[:2]
	fmt.Println("x:", x) // x = [0]

	y := a[2:]
	fmt.Println("y:", y) //y = [2 3]

	fmt.Println("================================")
	x = append(x, y...)
	fmt.Println("x:", x) // x = [0 2 3]
	//a = [0 2 3 3]
	fmt.Println("y:", y) // y = [3 3]
	//a = [0  2  3  3]

	x = append(x, y...)                // [0  2  3  3  3]
	fmt.Println("x:", x)               // [0  2  3  3  3]
	fmt.Println("a:", a, " ", "x:", x) // a = [0  2  3  3] // x = [0  2  3  3  3]
}
