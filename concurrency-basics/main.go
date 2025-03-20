package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello! ", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate slow operation, long-taking task

	fmt.Println("Hello! ", phrase)

	doneChan <- true // send signal to the channel, arrow points to the direction of data flow, in this case, from the goroutine to the channel

	close(doneChan) // close the channel after sending the signal
}

func main() {
	done := make(chan bool)
	// this willn't work as expected (if channel isn't used) because the main function will exit before the goroutines finish
	// so, we need to wait for the goroutines to finish before the main function exits
	go greet("hi", done)
	go greet("Running slow operation", done)

	// using channels to wait for goroutines to finish
	//here we are using a channel to wait for the goroutine to finish
	// the goroutine will send a signal to the channel when it finishes
	// and the main function will wait for the signal to be sent
	go slowGreet("Slow opeartion finished", done)
	go greet("How are you?", done)

	// isDone := <-done // waiting for data to come out of the channel, arrow points to the direction of data flow, in this case, from the channel to the main function

	for range done {

	}

}
