package main

// used to store memory address of a variable
// Before we start there are two important operators which we will use in pointers i.e.

//     * Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
//     & operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

func point() {
	var a *int
	var b int
	b = 10
	a = &b

	println(a)
	println(*a)

	z := add(a, b)
	println(z)
}

func add(x *int, y int) *int {
	// z := *x + y
	*x += y
	return x

}
