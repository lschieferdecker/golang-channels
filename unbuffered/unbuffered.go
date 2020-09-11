package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	unbufferedChan := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Calling asynchronously to demonstrate that the channel
	// will hold the flow until the receiver is ready.
	go sendDataToChan("test", unbufferedChan, wg)

	fmt.Printf("Sleeping...\n")
	time.Sleep(time.Second * 3)
	fmt.Printf("Waking up...\n")

	// Receiving from channel
	fmt.Printf("Received from channel: %v\n", <-unbufferedChan)

	wg.Wait()

}

func sendDataToChan(s string, unbufferedChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Sending to channel: %v\n", s)
	unbufferedChan <- s
	fmt.Printf("Sent to channel: %v\n", s)
}
