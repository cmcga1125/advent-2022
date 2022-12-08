package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Test")

	// array are fixed length
	var a [5]int
	a[4] = 100
	fmt.Println("set: ",a)

	// slices
	s := make([]string,0)
	fmt.Println("emp: ", s)

	fmt.Println("len: ", len(s))
	s = append(s,"Banana")
	fmt.Println("non-emp: ", s[0])

	fmt.Println("len: ", len(s))

	// maps
	m := make(map[string]int)
	m["apple"] = 7
	m["banana"] = 23
	fmt.Println("map: ", m)
	fmt.Println(m["banana"])
}