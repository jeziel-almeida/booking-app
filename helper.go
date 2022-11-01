package main

import (
	"fmt"
	"strings"
)

func validateUserInput(firstName, lastName, email string, userTickets uint) bool {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	isUserValid := true

	if !isValidName || !isValidEmail || !isValidTicketNumber {
		if !isValidName {
			fmt.Println("Your first or last name is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
		return !isUserValid
	}

	return isUserValid
}
