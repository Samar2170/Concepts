package main

import "fmt"

// func main() {
// 	// ...
// 	testStack()
// }

func testLinkedList() {
	ll := LinkedList{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)
	ll.InsertAt(1, 20)
	res := ll.SearchNode(5)
	fmt.Println(res)
	// ...
}

func testStack() {
	s := Stack{}
	s.Append(5)
	s.Append(10)
	s.Append(15)
	res := s.Pop()
	fmt.Println(res)
	fmt.Println(s)
	// ...
}
