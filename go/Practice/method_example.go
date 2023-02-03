package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (user User) getStatus() {
	fmt.Println("Status of", user.Name, " : ", user.Status)
}

func (user User) newEmail() {
	user.Email = "surajabc@abc.com"
	// email will only change here, but it will not effect actual object
	fmt.Println("Email of", user.Name, " : ", user.Email)
}

func main() {
	Suraj := User{"Suraj", 99, "surajxyz@xyz.com", true}
	fmt.Println("Suraj's details: ", Suraj)

	// get status using method
	Suraj.getStatus()

	// If you want to modify original mail, then we've to use
	// pointer, otherwise copy of an object will pass
	Suraj.newEmail()

	// if you print this, you can see that old mail will be there.
	fmt.Println("Suraj's details: ", Suraj)
}
