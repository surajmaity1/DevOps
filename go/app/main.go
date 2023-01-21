package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to", conferenceName, "booking app")
	fmt.Println("We've total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter Your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no. of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)

}
