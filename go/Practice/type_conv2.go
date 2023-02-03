// Convert string to float in go lang
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")
	input, _ := read.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Something Went Wrong!")
	} else {
		fmt.Println("You entered with adding one : ", num+1)
	}

}
