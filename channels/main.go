/**
Unbuffered channels: When you send data, it waits until someone is ready to receive it, hence blockings
-------------------------------------------------------------------------------------------------------
Buffered channels: Buffered channels allow sending and receiving without blocking until the buffer is full.
**/

package main

import (
	"fmt"
	"time"
)

// unbuffered channel example

func main() {
	ch := make(chan string)

	// Goroutine sending data
	go func() {

		time.Sleep(8 * time.Second) // mock block operation
		ch <- "Hello from goroutine"
	}()

	// Receiving data from channel, block here until data arrives
	msg := <-ch

	fmt.Println(msg)
}
