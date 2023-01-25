package main

import (
	"fmt"
	"sync"
	"time"
)

const doctorAppointmentTickets int = 100

var appointmentName = "DOCTOR APPOINTMENT"
var remainingAppointmentTickets uint = 100
var data = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	totalTickets uint
}

var wait_group = sync.WaitGroup{}

func main() {

	welcome_message()

	firstName, lastName, email, totalTickets := takeData()
	// Validation check in userInput.go by verifyInputData() function
	isValidName, isValidEmail, isValidTicketNumber := verifyInputData(firstName, lastName, email, totalTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookYourTickets(totalTickets, firstName, lastName, email)

		wait_group.Add(1)
		go sendTicketToMail(totalTickets, firstName, lastName, email)

		firstNames := usersFirstName()
		fmt.Printf("%v already booked tickets\n", firstNames)

		if remainingAppointmentTickets == 0 {
			fmt.Println("No tickets avaiable.\nBook another tickets next day")
		}
	} else {
		if !isValidName {
			fmt.Println("First name or Last name's length should be greater than two letters")
		}
		if !isValidEmail {
			fmt.Println("Invalid Email")
		}
		if !isValidTicketNumber {
			fmt.Println("no. of tickets should be > 1 or either tickets not available")
		}
	}
	wait_group.Wait()
}

func welcome_message() {
	fmt.Printf("Welcome To %v Booking App\n", appointmentName)
	fmt.Printf("TOTAL TICKETS: %v tickets.\nAVAILABLE TICKETS: %v.\n", doctorAppointmentTickets, remainingAppointmentTickets)
	fmt.Println("Book Tickets Now!")
}

func usersFirstName() []string {
	firstNames := []string{}
	for _, booking := range data {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func takeData() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter Email: ")
	fmt.Scan(&email)

	fmt.Println("Enter no. of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookYourTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingAppointmentTickets = remainingAppointmentTickets - userTickets

	var userData = UserData{
		firstName:    firstName,
		lastName:     lastName,
		email:        email,
		totalTickets: userTickets,
	}

	data = append(data, userData)
	fmt.Printf("\n\n\n\n\nTOTAL BOOKED DATA: \n%v\n\n\n\n", data)
	fmt.Println("You booked tickets successfully.")
	fmt.Println("You'll receieve confirmation mail sortly.")
	fmt.Printf("REMAINING TICKETS: %v\n", remainingAppointmentTickets)
}

func sendTicketToMail(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n\n\n_________________________")
	fmt.Printf("Sending tickets:  ...\n%v \nto email address %v\n", ticket, email)
	fmt.Println("_________________________")
	wait_group.Done()
}
