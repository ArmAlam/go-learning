package main

import "fmt"

func main() {

	var productNames [4]string = [4]string{"laptop", "mouse", "keyboard", "monitor"}

	// Declare a list of 4 float64 values, known as an array
	prices := [4]float64{10.0, 20.0, 30.0, 40.0}

	// access the 1st element of the list
	fmt.Println("1st price ", prices[0])

	// initialize the 1st element of the list
	productNames[0] = "tablet"
	fmt.Println("1st product ", productNames[0])

	// slice: a subset of an array
	// sampleSlice := productNames[:3] // slice from beginning to 3rd index, excluding 3rd index
	// sampleSlice := productNames[1:] // slice from beginning to last index

	sampleSlice := productNames[1:3] // slice from 1st to 3rd index, excluding 3rd index

	fmt.Println("sampleSlice ", sampleSlice) //[mouse keyboard]
}
