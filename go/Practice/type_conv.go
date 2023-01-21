// Type Conversion in Go
package main

import "fmt"

func main() {
	var a int = 32
	var b int = 4

	//var d int64 = b	// it will give compile time error as we
	// are performing an assignment between
	// mixed types i.e. int64 as int type

	var c float32 = float32(a) / float32(b)
	fmt.Println(c)
}
