package main

import (
	"fmt"
)

func evenOddChannel() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	numCh := make(chan int)

	go findEvenOdd(numbers, numCh)

	for num := range numCh {
		if num%2 == 0 {
			fmt.Println("Even Num ", num)
		} else {
			fmt.Println("Odd Num ", num)
		}
	}

	/*
				//MARK: 2
		        evenCh := make(chan int)
				oddCh := make(chan int)

				go findEven(numbers, evenCh)
				go findOdd(numbers, oddCh)

				fmt.Println("Even numbers:")
				for even := range evenCh {
					fmt.Println(even)
				}

				fmt.Println("\nOdd numbers:")
				for odd := range oddCh {
					fmt.Print(odd, " ")
				}
	*/

}

func findEvenOdd(arr []int, c chan int) {
	for _, num := range arr {
		if num%2 == 0 {
			c <- num

		} else {
			c <- num
		}
	}
	close(c)
}

// MARK: 2
func findEven(arr []int, evenChan chan int) {
	for _, num := range arr {
		if num%2 == 0 {
			evenChan <- num
			//fmt.Println("EvenNum:", c)
		}
	}
	close(evenChan)
}

func findOdd(arr []int, oddChan chan int) {
	for _, num := range arr {
		if num%2 != 0 {
			oddChan <- num
			// fmt.Println("OddNum:", c)
		}
	}
	close(oddChan)
}

// MARK: 1
func evenOddChannels() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	evenCh := make(chan int)
	oddCh := make(chan int)

	//MARK: 1
	go func() {
		for _, num := range numbers {
			evenCh <- num
		}
		close(evenCh)
	}()

	for _, num := range numbers {
		evenNum(num, evenCh)
	}

	go func() {
		for _, num := range numbers {
			oddCh <- num
		}
		close(oddCh)
	}()

	for val := range oddCh {
		oddNum(val, oddCh)
	}

}

func evenNum(num int, c chan int) {
	if num%2 == 0 {
		fmt.Println("EvenNum:", num)
	}
}

func oddNum(num int, c chan int) {
	if num%2 != 0 {
		fmt.Println("OddNum:", num)
	}
}
