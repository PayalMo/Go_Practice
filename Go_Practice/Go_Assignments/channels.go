package main

import "fmt"

func main8() {

	c := make(chan int)

	go addition(5, 10, c)
	add := <-c // receive from c
	fmt.Println("Addition is:", add)

	go sumOfSlice([]int{1, 2, 3, 4, 5}, c)
	sum := <-c // receive from c
	fmt.Println("Sum of slice is:", sum)

	bufferedChannel()
	unbufferedChannel()

}

func addition(a int, b int, c chan int) {
	c <- a + b // send sum to
	// c
}

func sumOfSlice(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// MARK: Buffereed Channel/Async
func bufferedChannel() {
	ch := make(chan int, 2) // buffered channel with capacity 2

	ch <- 1
	fmt.Println("Sent 1")

	ch <- 2
	fmt.Println("Sent 2")

	// The channel is full now, next send would block until a receive happens
	go func() {
		ch <- 3
		fmt.Println("Sent 3 (after space was freed)")
	}()

	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
	fmt.Println("Received:", <-ch)
}

// MARK: Unbuffered Channel/Sync
func unbufferedChannel() {
	ch := make(chan string) // unbuffered channel

	go func() {
		fmt.Println("Sending message...")
		ch <- "Hello from goroutine" // blocks until main receives
		fmt.Println("Message sent")
	}()

	msg := <-ch // waits until goroutine sends
	fmt.Println("Received:", msg)
}
