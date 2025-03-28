/**
Unbuffered channels: When you send data, it waits until someone is ready to receive it, hence blockings
-------------------------------------------------------------------------------------------------------
Buffered channels: Buffered channels allow sending and receiving without blocking until the buffer is full.
**/

package main

import (
	"fmt"
)

func main() {

	// ------------------- UNBUFFERED CHANNEL

	// ch := make(chan string)

	// // Goroutine sending data
	// go func() {

	// 	time.Sleep(8 * time.Second) // mock block operation
	// 	ch <- "Hello from goroutine"
	// }()

	// // Receiving data from channel, block here until data arrives
	// msg := <-ch

	// fmt.Println(msg)

	// ------------------- BUFFERED CHANNEL
	ch := make(chan string, 2) // buffered channel with a capacity of 2

	go func() {
		ch <- "Message 1"
		ch <- "Message 2"
	}()

	// Receiving data from channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
