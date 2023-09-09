package main

import "fmt"

func generator(msg string) <-chan string { // returns a recieve only channel
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func testGenerator() {
	c := generator("boring!")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
}

func main() {
	testFanIn()
}

func serviceLevelChannel() {
	joe := generator("Joe")
	ann := generator("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(i1, i2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-i1
		}
	}()
	go func() {
		for {
			c <- <-i2
		}
	}()
	return c
}

func testFanIn() {
	c := fanIn(generator("joe"), generator("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

}
