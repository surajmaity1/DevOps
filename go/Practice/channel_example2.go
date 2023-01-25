// Channel in Golang
package main

import "fmt"
func justPrint(c chan int) {
	for i := 0; i < 6; i++ {
		c <- i
	}
	close(c)
}
func main() {
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


	fmt.Println("\n\nUSE FOR LOOP IN THE CHANNEL:")
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

	// Finding the length of the channel
	channel5 := make(chan int, 6)

	channel5 <- 1
	channel5 <- 2
	channel5 <- 3
	channel5 <- 4
	channel5 <- 5
	fmt.Println("channel5's length:",len(channel5), "& capacity:", cap(channel5))

	
}
