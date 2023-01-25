// Slice in Go
package main

import (
	"fmt"
	"reflect"
	"sort"
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

	fmt.Println("arr's type:", reflect.ValueOf(arr).Kind())
	fmt.Println("slice1's type:", reflect.ValueOf(slice1).Kind())

	arr2 := [7]string{"Hey", ",", "I", "am", "learning", "Go", "Language"}

	//array_name[low:high]
	slice2 := arr2[:5]
	fmt.Println("slice2's type:", reflect.ValueOf(slice2).Kind())
	fmt.Println(slice2)
	fmt.Println("slice2's length: ", len(slice2), "capacity:", cap(slice2))

	// Another way to create slice
	var slice3 = []string{"a", "b", "c"}
	slice3 = append(slice3, "d")
	slice3 = append(slice3, "e")
	fmt.Println("slice3's len:", len(slice3), ", cap:", cap(slice3))

	//make() function
	var slice4 = make([]int, 4, 5)
	fmt.Printf("Slice4 = %v, \nlength = %d, \ncapacity = %d\n", slice4, len(slice4), cap(slice4))

	// iterate over slice using loop

	for i := 0; i < len(slice4); i++ {
		fmt.Println(slice4[i])
	}

	// Iterate slice
	// using range in for loop
	for index, element := range slice4 {
		fmt.Println("slice4[", index, "]=", element)
	}

	// Iterate slice
	// using range in for loop
	// without index
	for _, element := range slice4 {
		fmt.Printf("Element = %v\n", element)
	}

	// modify array using slice
	var arr3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Before modify: ", arr3)
	slice5 := arr3[:]

	slice5[2] = 300
	slice5[4] = 500
	fmt.Println("After modifying using slice: ", arr3)

	// Comparison of Slice
	fmt.Println("slice5 is null?:", slice5 == nil)

	// note: you can't compare two slice in go
	//only you can compare with one slice with null
	//fmt.Println("slice5 = slice3?:", slice5 == slice3) compile time error

	// Note: If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function.

	fmt.Println("slice5 = slice3?:", reflect.DeepEqual(slice5, slice3))

	// multi-dimensional slice

	/*
		var slice6 = [][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		}
		above multi-dime is same as below
	*/
	var slice6 = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(slice6)

	// SORT SLICE
	slice7 := []int{43, 2, 12, 444, 32}
	sort.Ints(slice7)
	fmt.Println("sorted slice:", slice7)
}
