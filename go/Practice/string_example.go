package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main(){
	str := "Hello World!"

	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))
	fmt.Println(len(str))
	fmt.Println("Ascii value of A is ","Aabc"[2])
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
}
