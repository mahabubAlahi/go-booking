package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//book Tickets
			bookTickets(remainingTickets, userTickets, firstName, lastName, emailAddress)

			//Print First Names

			firstNames := getFirstNames()

			fmt.Printf("All bookings first name are: %v\n", firstNames)

			if remainingTickets == 0 {
				break
			}
		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email addres you entered does not contain @ sign")
			}
			if !isValidTicketNumber {

				fmt.Println("number of tickets you entered is wrong")
			}
		}

	}

}

func greetUsers() {
	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have tickets ", conferenceTickets, "and available tickets ", remainingTickets)
	fmt.Println("Get your tickets here")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email Address: ")
	fmt.Scan(&emailAddress)

	fmt.Println("Enter how many tickets you want: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, emailAddress, userTickets
}

func bookTickets(remainingTickets int, userTickets int, firstName string, lastName string, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	//Create a booking map

	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["emailAddress"] = emailAddress
	userData["numberOfTickets"] = strconv.Itoa(userTickets)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
