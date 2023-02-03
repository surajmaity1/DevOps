package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome")

	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter value: ")
	input, err := inputReader.ReadString('\n')

	fmt.Println("You entered:", input)
	fmt.Printf("Type: %T,\nError: %v\n", input, err)
}
