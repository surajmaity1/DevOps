// Slice in Go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var slice1 []string
	slice1 = append(slice1, "Item 123")
	slice1 = append(slice1, "Item 434")
	slice1 = append(slice1, "Item 4564")
	slice1 = append(slice1, "Item 7656")

	fmt.Println("slice2's type:", reflect.ValueOf(slice1).Kind())
	fmt.Println("Length of slice", len(slice1))
	fmt.Println("Item at index 2:", slice1[1])

	arr := [3]string{"Item1", "Item2", "Item3"}
	//slice2 := arr

	fmt.Println("arr's type:", reflect.ValueOf(arr).Kind())
	fmt.Println("slice2's type:", reflect.ValueOf(slice1).Kind())
}
