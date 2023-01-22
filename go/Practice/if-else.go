// If-else in go
package main

import "fmt"

func printNote(str string){
	fmt.Println(str)
}


func main() {
	
	var val int
	fmt.Print("Enter a number: ")
	fmt.Scan(&val)

	if val < 10 && val > 0 {
		printNote("value entered less than 10 && greater than 0")
	} else if val > 100 && val < 1000 {
		printNote("value entered less than 1000 && greater than 100")
	} else {
		fmt.Println("It is:", val)
	}
}
