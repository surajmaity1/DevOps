// Different Ways to Find the Type of Variable in Golang
package main

import (
	"fmt"
	"reflect"
)

func main() {

	arr := [3]string{"Item1", "Item2", "Item3"}

	fmt.Println("arr's type:", reflect.TypeOf(arr))
	fmt.Printf("arr's type: %T \n", arr)
	fmt.Println("arr's type:", reflect.ValueOf(arr).Kind())
}
