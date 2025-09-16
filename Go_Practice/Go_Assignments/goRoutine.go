package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main7() {

	for _, link := range links {
		wg.Add(1) // Increase counter for each goroutine
		go printLink(link, &wg)
	}

	for _, color := range colors {
		wg.Add(1) // Increase counter for each goroutine
		go printColors(color, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish

	fmt.Println("All links printed")
	fmt.Println("All colors printed")
	fmt.Println("--------End of All Tasks-----")
}

// MARK: 1
var links = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.twitter.com",
	"https://www.github.com",
}

func printLink(link string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	fmt.Println("Printed link:", link)

}

// MARK: 2
var colors = []string{
	"Red",
	"Blue",
	"Green",
	"Yellow",
}

func printColors(color string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement counter when done
	fmt.Println("Colors:", color)
}
