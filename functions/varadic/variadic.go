package main

import "fmt"

func main() {

	numbers := []int{10, 20, 30, 40, 50}

	numResult := sumup(numbers...)
	result := sumup(1, 2, 3, 4, 5)

	fmt.Println(numResult)
	fmt.Println(result)
}

// can receive any number of integers, this is known as variadic function
func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
