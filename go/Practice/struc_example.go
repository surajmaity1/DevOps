package main

import "fmt"

type Student struct{
	/*
	Name string
	Age int
	blood_grp string

	or below as it is
	*/
	Name, blood_grp string
	Age int
}

func main(){
	s1 := Student{"Suraj", "X", 21}
	fmt.Println("s1:", s1)

	s2 := Student{Name: "Sudipta"}
	fmt.Println("s2:", s2)

	var s3 Student
	fmt.Println("s3:", s3)


	// access fields of struct
	fmt.Println("s2's name:", s2.Name)


	// THE POINTER TO STRUCT
	struct_ptr1 := &Student{Name: "Amit"}
	fmt.Println("struct_ptr1's name:", (*struct_ptr1).Name)

	// The Golang gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field.
	fmt.Println("struct_ptr1's name:", struct_ptr1.Name)

}