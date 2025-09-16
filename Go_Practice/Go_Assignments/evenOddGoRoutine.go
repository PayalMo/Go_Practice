package main

import (
	"fmt"
	"sync"
)

func main9() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	for _, num := range numbers {
		wg.Add(2) // Increase counter for each pair of goroutines
		go evenGoRoutine(num, &wg)
		go oddGoRoutine(num, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All even and odd numbers printed")
	fmt.Println("--------End of All Tasks-----")
}

func evenGoRoutine(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	if num%2 == 0 {
		fmt.Println("Even Number:", num)
	}
}

func oddGoRoutine(num int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	if num%2 != 0 {
		fmt.Println("Odd Number:", num)
	}
}
