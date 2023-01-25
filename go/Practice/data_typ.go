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

	var x uint32 = 43 // Using 32-bit unsigned int
	fmt.Println(x, x-32)

	var a int64 = 324 // Using 64-bit signed int
	fmt.Println(a+323, a*2)

	var b float64 = 42342.2
	fmt.Println(b)

	var c complex64 = complex(43, 2)   //Complex numbers which contain float32 as a real and imaginary component.
	var d complex128 = complex(333, 3) //Complex numbers which contain float64 as a real and imaginary component.

	fmt.Println(c, d)

	var e = true
	fmt.Printf("e: %v and types of e: %T\n", e, e)

	str1 := "Hello"
	fmt.Println(str1 == "Hello")
}
