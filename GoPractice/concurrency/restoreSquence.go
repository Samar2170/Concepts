package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func testRSequencing() {
	c := make(chan Message)
	waitForIt := make(chan bool)

	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s: %d", "msg", i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
}

func main() {
	testRSequencing()
}
