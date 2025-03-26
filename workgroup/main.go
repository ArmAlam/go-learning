package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func processFile(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", fileName, err)
		return
	}
	fmt.Printf("Processed %s (%d bytes) in %v content: %s \n", fileName, len(data), time.Since(start), string(data))
}

func main() {
	files := []string{"log1.txt", "log2.txt", "log3.txt"}
	var wg sync.WaitGroup

	start := time.Now()
	for _, file := range files {
		wg.Add(1)
		go processFile(file, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Printf("Total time: %v\n", time.Since(start))
}
