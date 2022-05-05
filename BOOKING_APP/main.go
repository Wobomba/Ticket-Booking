package main

import (
	"fmt"
	"strconv"
)

func main(){
	//same as var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
    const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available\n", conferenceTickets,remainingTickets)
	fmt.Print("Get your tickets here to attend\n")

	var userName string
	var userTickets int

	/*var bookings = [50]string{"Nana", "Nicole", "Peter"} --> using arrays
	var bookings [50]string
	bookings[0] = userName + " " + strconv.Itoa(userTickets)
	fmt.Printf("List of attendees: %v\n", bookings)
	fmt.Printf("The first attendee value: %v\n", bookings[0])*/

	//creating a slice
	bookings := []string{}
	for {
		//username input
		fmt.Println("Enter your name: ")
		fmt.Scan(&userName)
		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)
		bookings = append(bookings, userName+"   "+strconv.Itoa(userTickets))
		fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
		fmt.Printf("The slice version: %v\n", bookings)
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)
	}
	

}
