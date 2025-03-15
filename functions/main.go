package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	transformed := transformNumbers(&numbers, func(number int) int { // pass an anonymous function
		return number * 2
	})

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}
