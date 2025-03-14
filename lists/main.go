package main

import "fmt"

// type aliases
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// make made code more efficient
	usernames := make([]string, 2, 5) // 2  -> size of the slice, 5 -> capacity of the slice, helps to reduce memory allocation

	usernames[0] = "arman"

	usernames = append(usernames, "zarman")
	usernames = append(usernames, "barman")

	sameMaps := make(floatMap, 2) // 2 -> size of the map
	sameMaps["sample1"] = 1.0
	sameMaps["sample2"] = 2.0 // go will use existing map size
	sameMaps["sample3"] = 3.0 // go will resize the map automatically

	sameMaps.output()

	// loop for slice
	for index, value := range usernames {
		fmt.Println(index, value)
	}

	// loop for maps
	for key, value := range sameMaps {
		fmt.Println(key, value)
	}
}
