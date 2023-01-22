package main

import "fmt"

// Function without return type
func fun1(str1 string, str2 string){
	fmt.Println("You said:", str1,str2)
}
// Function with return type

func fun2(str1 string, str2 string)string{
	return "You said: " + str1 + " " + str2
}

// Function to return multi return value
func fun3(str1 string, str2 string) (string, string){
	str1 += " ready to go"
	str2 += " ready to go"
	return str1, str2
}

// Function to swap two number using call by value
func swapCall(a int, b int){
	a = a ^ b
	b = a ^ b
	a = a ^ b
}

// Function to swap two number using call by reference
func swapRef(a *int, b *int){
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main(){
	fun1("hello", "everyone")
	fmt.Println(fun2("Good", "day"))

	i_am, we_are := fun3("I am", "We are")
	fmt.Println(i_am, "\n", we_are)

	var a = 5
	var b = 6
	fmt.Println("Before swapping : a:", a, ", b:", b)
	swapCall(a, b)
	fmt.Println("After swapping using call by val: a:", a, ", b:", b)


	fmt.Println("Before swapping: a:", a, ", b:", b)
	swapRef(&a, &b)
	fmt.Println("After swappingusing call by ref: a:", a, ", b:", b)

}