package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printNote(c int) {
	fmt.Println("It is :", c)
}
func checkSwitch(val int) {
	switch val {
	case 1:
		printNote(1)
	case 2:
		printNote(2)
		fallthrough
	case 3:
		printNote(3)
	case 4:
		printNote(4)
		fallthrough
	case 5:
		printNote(5)
	default:
		fmt.Print("Not match\n")
	}
}

func main() {
	// IT WILL GENERATE RANDOM VALUE BETWEEN 0 AND 5
	rand.Seed(time.Now().Unix())
	var val = rand.Intn(6)
	fmt.Println("val :", val)
	checkSwitch(val)

	switch m := 8; m {
	case 1:
		printNote(1)
	case 2:
		printNote(2)
	case 3:
		printNote(3)
	case 4, 8:
		if val == 4 {
			printNote(4)
		} else {
			printNote(8)
		}
	case 5:
		printNote(5)
	default:
		fmt.Print("Not match\n")
	}

}
