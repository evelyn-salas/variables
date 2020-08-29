package main

import (
	"fmt"
)

func main() {
	//basic print command
	fmt.Printf("hello, world!!\n")

	//declaring variables, method 1
	//good for initializing later
	//can be used to declare variables
	//outside of the package
	var i int
	i = 42
	i = 27
	fmt.Println(i)

	//declaring variables, method 2
	//good if go gets the type wrong
	var j int = 15
	fmt.Println(j)

	//decalring variables, method 3
	k := 14.
	fmt.Println(k)

	//how to check the variable type
	fmt.Printf("%v, %T", k, k)

	//type conversions
	var m float32
	m = float32(i)
	fmt.Printf("%v, %T", m, m)
}
