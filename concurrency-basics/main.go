package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello! ", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second) // simulate slow operation, long-taking task

	fmt.Println("Hello! ", phrase)
}

func main() {
	greet("hi")
	greet("Running slow operation")
	slowGreet("Slow opeartion finished")
	greet("How are you?")
}
