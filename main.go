package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, emailAddress, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//book Tickets
			remainingTickets, bookings := bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, emailAddress, conferenceName)

			//Print First Names

			firstNames := getFirstNames(bookings)

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

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		names := strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
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

func bookTickets(remainingTickets int, userTickets int, bookings []string, firstName string, lastName string, emailAddress string, conferenceName string) (int, []string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("These are all our bookings: %v\n", bookings)

	return remainingTickets, bookings
}
