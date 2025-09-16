package main

import (
	"fmt"
)

func main1() {
	numbers := [5]int{1, 2, 3, 4, 5}
	for _, value := range numbers {
		fmt.Printf("Value: %d\n", value)
	}
}
