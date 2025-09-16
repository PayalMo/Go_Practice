package main

import (
	"fmt"
)

const maxSpace = 20

var currNumberOfBlocks int

type VehicleParkingData interface {
	numberOfBlocks() int
}

func main6() {
	c := Car{}
	b := Bike{}

	s := []VehicleParkingData{c, b, c, b, b, c, c, c, b, c, b, b, b, c} // MARK: Slice of Interface

	for _, value := range s {
		dataInSlice := value.numberOfBlocks()
		println("Data in slice:", dataInSlice)
		currNumberOfBlocks += dataInSlice
		isParkingAvailable()
	}
}

func isParkingAvailable() {
	if currNumberOfBlocks > maxSpace {
		fmt.Println("We can not park as parking is full already")
	} else {
		fmt.Println("Yes you can park it as current space is ", currNumberOfBlocks)
	}
}

type Car struct{}

type Bike struct{}

func (c Car) numberOfBlocks() int {
	return 2
}

func (b Bike) numberOfBlocks() int {
	return 1
}

// func addNumberOfBlocks(blocks VehicleParkingData) []int {
// 	numOfBlocks := []int{}
// 	numOfBlocks = append(numOfBlocks, blocks.numberOfBlocks())
// 	return numOfBlocks
// }

/*
func main6() {
	c := Carss{}
	b := Bikess{}

	arr := []int{} // MARK: Array of Interface

	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, b.numberOfBlocks())
	arr = append(arr, c.numberOfBlocks())

	for _, value := range arr {

		println("Data in array:", value)
		currNumberOfBlocks += value
		isParkingAvailable()
	}

}
*/
