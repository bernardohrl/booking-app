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

	var userName string
	var userTickets int
	userName = "Tom"
	userTickets = 2
	fmt.Printf("\n%s bougth %d tickets\n", userName, userTickets)
}
