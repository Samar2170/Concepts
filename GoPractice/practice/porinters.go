package main

import "fmt"

func main() {
	// i := 1
	// j := &i
	// k := *j
	// l := &j
	// fmt.Println(i)
	// fmt.Println(j)
	// fmt.Println(k)
	// fmt.Println(l)

	var x *int
	y := 10
	x = &y
	fmt.Println(x)
	fmt.Println(*x)
}

// * is used to access value, & to access pointer
// *int, * used with type to declare pointer var, and used to access value
// & address of operator
