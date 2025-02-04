package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50

	fmt.Println("wellcome to", conferenceName, ". Total tickets:", conferenceTickets, ". Total available:", remainingTickets)
	fmt.Printf("wellcome to %s. Total tickets: %d. Total available: %d", conferenceName, conferenceTickets, remainingTickets)

	var bookings [50]string
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Print("\nEnter first name: ")
	fmt.Scan(&firstName)
	fmt.Print("\nEnter last name: ")
	fmt.Scan(&lastName)
	fmt.Print("How many tickets do you want? ")
	fmt.Scan(&userTickets)

	bookings[0] = firstName + " " + lastName
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("\n%s bougth %d tickets\n", bookings[0], userTickets)
	fmt.Printf("%d tickets remaining\n", remainingTickets)

	fmt.Printf("Array: %v\n", bookings)
	fmt.Printf("Array: %v\n", bookings[0])
	fmt.Printf("Array: %T\n", bookings)
	fmt.Printf("Array: %v\n", len(bookings))
}
