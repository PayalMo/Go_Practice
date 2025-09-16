package main

import (
	"fmt"
)

func numAndLetters() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26}
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	numCh := make(chan int)
	letterCh := make(chan string)

	go numbersData(numbers, numCh)
	go lettersData(letters, letterCh)

	pairCount := len(numbers[:26]) / 2
	for i := 0; i < pairCount; i++ {
		fmt.Printf("%d%d%s%s", <-numCh, <-numCh, <-letterCh, <-letterCh)
	}

}

func numbersData(numArr []int, numChan chan int) {
	for _, num := range numArr {
		numChan <- num
	}
	close(numChan)

}

func lettersData(lettersArr []string, letterChan chan string) {
	for _, letter := range lettersArr {
		letterChan <- letter
	}
	close(letterChan)

}

/*
var prevNum, nextNum int
	var prevLetter, nextLetter string

	pairCount := len(numbers[:26]) / 2
	for i := 0; i < pairCount; i++ {
		prevNum = <-numCh
		nextNum = <-numCh
		prevLetter = <-letterCh
		nextLetter = <-letterCh
		fmt.Printf("%d%d%s%s", prevNum, nextNum, prevLetter, nextLetter)
	}

*/
