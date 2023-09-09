package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         // points to i, so you access the pointer of a var using &
	fmt.Println(*p) //you access the value using *operator if its a pointer

	*p = 21 //set value of i via pointer
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
