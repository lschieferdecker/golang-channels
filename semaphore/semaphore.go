package main

import (
	"fmt"
	"time"
)

// Creating a new channel with capacity 5. It will
// be used to limit processing throughput.
var semaphore = make(chan int, 5)

func main() {
	// You may change this to an queue that receives requests
	// to be processed. Here I'm just simulating a thousand
	// elements to be processed.
	for i := 0; i <= 1000; i++ {
		go handle(i)
	}

	fmt.Printf("Done \n")
	// Prevent main from exiting
	fmt.Scanln()
}

func handle(i int) {
	semaphore <- 1
	process(i)
	<-semaphore
}

func process(i int) {
	fmt.Printf("Processing %v\n", i)
	time.Sleep(time.Second)
	fmt.Printf("Processed %v\n", i)
}
