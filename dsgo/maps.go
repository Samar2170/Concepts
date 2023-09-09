package main

import "fmt"

func main() {
	myMap := make(map[string]string)
	myMap["a"] = "a"

	myMap2 := map[string]string{
		"a": "a",
		"b": "b",
		"c": "c",
	}

	myMap2["d"] = "d"
	fmt.Println(myMap2)
}
