package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello! ", phrase)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate slow operation, long-taking task

	fmt.Println("Hello! ", phrase)

	doneChan <- true // send signal to the channel, arrow points to the direction of data flow, in this case, from the goroutine to the channel
}

func main() {
	// this willn't work as expected because the main function will exit before the goroutines finish
	// so, we need to wait for the goroutines to finish before the main function exits
	// go greet("hi")
	// go greet("Running slow operation")

	// using channels to wait for goroutines to finish
	//here we are using a channel to wait for the goroutine to finish
	// the goroutine will send a signal to the channel when it finishes
	// and the main function will wait for the signal to be sent
	done := make(chan bool)
	go slowGreet("Slow opeartion finished", done)
	// go greet("How are you?")

	isDone := <-done

	fmt.Println(isDone) // waiting for data to come out of the channel, arrow points to the direction of data flow, in this case, from the channel to the main function
}
