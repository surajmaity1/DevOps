// Loops in Go
package main

import (
	"fmt"
)

func main() {

	//Infinite loop
	for {
		fmt.Println("Hello Everyone :)")
		break //If we can't use break it will go infinity
	}

	// loop works
	for i := 1; i < 20; i++ {
		fmt.Print(i, " ")
	}

	//for loop as while Loop:
	j := 23
	for j < 101 {
		fmt.Print(j, " ")
		j += 3
	}
	fmt.Println()

	slice1 := []int{43, 2, 12, 444, 32}

	for index, element := range slice1 {
		fmt.Println("Index:", index, "Element:", element)
	}

	var str1 string = "Hey Man!"

	for index, element := range str1 {
		// note here element will print ascii value of char
		// always convert it to character
		fmt.Printf("Index: %v, Element:%v\n", index, string(element))
	}

	// Iterate through map
	// declare : map[key]value
	map1 := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	for key, val := range map1{
		fmt.Println("key:", key, "val:", val)
	}



	//use for loop using channel
	channel5 := make(chan int)
    go func(){
		channel5 <- 11
		channel5 <- 12
		channel5 <- 13
		channel5 <- 14
		close(channel5)
    }()
    for i:= range channel5 {
       fmt.Println(i) 
    }
}
