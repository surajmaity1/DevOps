package main

import "fmt"
var var1 = 30

func main() {
	/*
	Note: What will happen there exists a local variable with the same name as that of the global variable inside a function?

	The answer is simple i.e. compiler will give preference to the local variable.
	*/
	var var1 = 32
	fmt.Println("var1: ", var1)
}
