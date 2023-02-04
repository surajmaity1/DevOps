package main

import "fmt"

func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	/*
		stack ---> 	Three, Two, One		[it will og as LIFO]
	*/
	// note: non-defer line will print first and as it is
	fmt.Println("Hello")
	myDeferFun()
	fmt.Println()
	/*
		One, Two, Three
		note: here, lifo will be occur -->> it will print three, two and one
	*/
}

func myDeferFun() {

	/*
		stack ---> 9,8,7,6,5,4,3,2,1,0 		[it will og as LIFO]
	*/
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%v ", i)
	}
}
