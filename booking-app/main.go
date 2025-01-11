package main

import "fmt"

func main() {

	var conferenceName = "Go Conference"

	const conferenceTickets = 100

	fmt.Printf("Total Tickets %v Available\n", conferenceTickets)

	fmt.Printf("Welcome To %v Booking App\n", conferenceName)

	var userName string
	var userTicket int
	fmt.Print("Enter Your Name: ")
	fmt.Scan(&userName) //& is used to get the address of the variable. It is pointer to the variable address in memmory

	fmt.Print("Enter Number of Tickets: ")
	fmt.Scan(&userTicket)

	fmt.Printf("Thank You for %v booking %v tickets", userName, userTicket)

}
