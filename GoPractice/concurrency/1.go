package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go is latest on the Newsqueak-Alef-Limbo family tree, distinguised by first class channels
//  Rough analogy: writing to a file by name (process, Erlang) vs. writing to a file descriptor (channel, Go).

func main() {
	simple()

}
func simple() {
	c := make(chan string)
	go boring("boring!", &c) // this thread runs on forever, below thread exits after 5 loops
	//
	for i := 0; i < 5; i++ {

		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

// <- c will wait for the sender to send a value
// -> c will wait for reciever to recieve a value
// Thus channels both communicate and synchronize.
// Unsychrnoized channels will be buffered channels, i.e for live streaming
func boring(msg string, c *chan string) {
	for i := 0; ; i++ {
		*c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

// Goroutine
// independenly executing function, launched by go statement
// has its own call stack, which grows and shrinks as required
// very cheap, practical to have thousands, even hundreds of thousands of goroutines
// not a thread, not an OS thread, managed by Go runtime
// goroutines are multiplexed dynamically onto threads as needed to keep all running
