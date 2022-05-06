package main

import (
	"fmt"
	"strconv"
	"strings"
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

	//creating a slice and making it dynamic
	bookings := []string{}
	for {
		//username input and validation
		fmt.Println("Enter your name: ")
		fmt.Scan(&userName)
		fmt.Println("Enter the number of tickets: ")
		fmt.Scan(&userTickets)
		
		isValidName := len(userName) >= 2 
		isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets

		if isValidName || isValidTicketNum{
			bookings = append(bookings, userName+"   "+strconv.Itoa(userTickets))
			fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
			fmt.Printf("The slice version: %v\n", bookings)
			//if the user inputs more tickets than the remaining tickets
			if userTickets <= remainingTickets{
				remainingTickets = remainingTickets - userTickets
				fmt.Printf("%v tickets remaining for %v\n", remainingTickets,conferenceName)
				
				//iterating over a list
				userNames := []string{}
				for _, booking :=range bookings{
					var names = strings.Fields(booking) 
					userNames = append(userNames, names[0])

				}
				fmt.Printf("The names of the bookings are: %v\n", userNames)

				noTicketLeft := remainingTickets == 0
				if noTicketLeft{
					fmt.Println("Our conference is booked out. Come back next year")
					break
				}
			} else {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
				continue
			}

		}else{
			//specifying what input data is wrong
			if !isValidName{
				fmt.Println("Your user name is too short")
			}
			if !isValidTicketNum{
				fmt.Println("The number of tickets entered is not valid")
			}
			
		}
	}
	

}
