package refursion

import "fmt"

func main() {
	fact := factorial(10)

	fmt.Println(fact)
}

func factorial(number int) int {

	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}
