package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	bookings := []string{}

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have tickets ", conferenceTickets, "and available tickets ", remainingTickets)
	fmt.Println("Get your tickets here")

	for {
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

		remainingTickets = remainingTickets - userTickets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("These are all our bookings: %v\n", bookings)

		fmt.Printf("All bookings first name are: %v\n", firstNames)
	}

}
