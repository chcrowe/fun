package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

// The `fast` and `slow` functions do the same thing
// but `slow` takes more time to complete
func fast(num int, out chan<- int) {
	result := num * 2
	time.Sleep(5 * time.Millisecond)
	out <- result

}

func slow(num int, out chan<- int) {
	result := num * 2
	time.Sleep(15 * time.Millisecond)
	out <- result
}

func signalCtrlC_handler(done chan bool) {

	sig := make(chan os.Signal, 1) //signal channel to capture CTRL+C
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Print(" shutting down ")
	for j := 0; j < 3; j++ {
		time.Sleep(time.Second * 1)
		fmt.Print(".")
	}
	fmt.Println(" [done]")

	done <- true
}

func main2() {
	out1 := make(chan int)
	out2 := make(chan int)

	done := make(chan bool, 1)   //this is the main service channel
	go signalCtrlC_handler(done) //register the CTRL+C handler

	// we start both fast and slow in different
	// goroutines with different channels
	go fast(2, out1)
	go slow(3, out2)

	for {
		// perform some action depending on which channel
		// receives information first
		select {
		case res := <-out1:
			fmt.Println("fast finished first, result:", res)
		case res := <-out2:
			fmt.Println("slow finished first, result:", res)
		case res := <-done:
			fmt.Println("os signal received, result:", res)
			os.Exit(0)
		}
	}
}
