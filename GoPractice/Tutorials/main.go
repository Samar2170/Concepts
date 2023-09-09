package main

import "fmt"

func main() {
	// testCh()
	// point()
	// goTest()
	testConcurrency()
}
func goTest() {
	for i := 0; i < 50; i++ {
		go func() {
			fmt.Println("hello")
		}()
	}
}
