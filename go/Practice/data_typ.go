// DATA TYPES IN Go

package main

import "fmt"

func main() {
	var name = "Suraj"
	var age int = 22

	var hobby string
	hobby = "Male"

	// print variable
	fmt.Printf("Name: %v, age: %v and hobby: %v\n", name, age, hobby)

	// check type
	fmt.Printf("Name: %T, age: %T and hobby: %T\n", name, age, hobby)
}
