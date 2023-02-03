package main

import (
	"fmt"
	"time"
)
func main(){
	/*
	fmt.Println()

	chan1 := make(chan string, 2)

	go func ()  {
		time.Sleep(4 * time.Second)
		chan1 <- "hello in chan1"
	}()
	
	select {
	case output_chan1 := <- chan1 :
		fmt.Println(output_chan1)
		
	case <- time.After(5 * time.Second) :
		fmt.Println("Printing this line after 5 second for chan1")
	}


	chan2 := make(chan string, 2)

	go func ()  {
		time.Sleep(8 * time.Second)
		chan2 <- "hello in chan1"
	}()
	
	select {
	case output_chan2 := <- chan2 :
		fmt.Println(output_chan2)
		
	case <- time.After(7 * time.Second) :
		fmt.Println("Printing this line after 7 second for chan2")
	}

	*/
	// FETCH TIME AND MAKE IT A FORMATTED OUTPUT

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2023, time.February, 3, 11, 8, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}