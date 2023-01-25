package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Map Example")

	// Define a map using make() fun
	var map1 = make(map[string]string)
	map1["1"] = "one"
	fmt.Println("map1: ", map1)
	fmt.Println("map1: ", reflect.TypeOf(map1))

	// Define a map without make() fun
	map2 := map[int]string{
		31 : "Thirty One",
	}
	map2[32] = "Thirty Two"
	fmt.Println("map2: ", map2)

	fmt.Println("map2: ", reflect.TypeOf(map2))

	// Define a map without make() fun
	// Note: map3 is a nil
	var map3 map[float64]bool
	fmt.Println(reflect.TypeOf(map3))

	if map3 == nil {
		fmt.Println(true)
	}

	// DEFINING A LIST OF MAPS
	var listOfMaps = make([]map[int]string, 0)
	fmt.Println(reflect.TypeOf(listOfMaps))




	// check the key is available or not
	_, check := map2[23]
	fmt.Println("23 key is present in map2?:", check)

	// delete key 
	fmt.Println("Original map2:", map2)
	delete(map2, 31)
	fmt.Println("After deleting 31 in map2:", map2)
}
