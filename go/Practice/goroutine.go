package main

import (
	"fmt"
	"time"
)

func main() {
	go caller("Hey")
	caller("Hi")
}

func caller(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		fmt.Println(s)
	}
}
