// array in Go
package main

import "fmt"

func main() {
	var arr1 [5]int
	var arr2 [5]string
	arr3 := [5]string{"a", "b", "c", "d", "e"}

	fmt.Printf("arr1: %v\n", arr1)
	fmt.Printf("arr2: %v\n", arr2)
	fmt.Printf("arr3: %v\n", arr3)

	for i := 0; i < len(arr1); i++ {
		fmt.Print(" ", arr1[i])
	}
	fmt.Println()
	arr2d := [2][2]string{{"a", "b"}, {"d", "c"}}

	for i := 0; i < len(arr2d); i++ {
		fmt.Println(arr2d[i])
	}

	//In an array, if ellipsis ‘‘…’’ become visible at the place of length,
	//then the length of the array is determined by the initialized elements.
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 45, 7, 8, 9, 3, 4, 4, 5, 23}
	fmt.Println(arr4)

	//In Go language, an array is of value type not of reference type.

	/*
		So when the array is assigned to a new variable,
		then the changes made in the new variable do not affect the original array.

	*/
	arr5 := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(arr5)

	arr6 := arr5
	arr6[0] = "first index changed"
	fmt.Println(arr5)
	fmt.Println(arr6)

	//we can directly compare two arrays using == operator.
	fmt.Printf("arr3 == arr5?: %v\n", arr3 == arr5)
	fmt.Printf("arr3 == arr6?: %v\n", arr3 == arr6)
}
