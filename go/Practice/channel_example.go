// Channel in Golang
package main

import "fmt"

func printFunc(channel1 chan int) {
	fmt.Println(234 + <-channel1)
}

func justPrint(c chan int) {
	for i := 0; i < 6; i++ {
		c <- i
	}
	close(c)
}
func main() {
	
	// create a channel using var keyword
	var channel1 chan int
	fmt.Printf("channel1's val: %v,  type: %T\n", channel1, channel1)

	// create a channel using make() fun
	channel2 := make(chan int)
	fmt.Printf("channel2's val: %v,  type: %T\n", channel2, channel2)

	//Send and Receive Data From a Channel
	go printFunc(channel1)
	channel1 <- 2
	

	// run this one, after commenting above codes
	// Closing a Channel
	fmt.Println("------------")
	channel3 := make(chan int)
	go justPrint(channel3)

	for {
		result, check := <-channel3

		if check == false {
			fmt.Println("Channel close", check)
			break
		}
		fmt.Println("Channel open", result, check)
	}

	
	// HOW TO USE FOR LOOP IN THE CHANNEL
	channel4 := make(chan int)

	go func ()  {
		channel4 <- 1
		channel4 <- 2
		channel4 <- 3
		channel4 <- 4
		channel4 <- 5
		close(channel4)
	}()

	for data := range channel4 {
		fmt.Println(data)
	}
}
