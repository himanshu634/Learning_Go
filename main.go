package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets int = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName  string
	email     string
	tickets   int
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()

	isValidTickets, isValidEmail, isValidName := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTickets {

		bookTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("First names are %v\n", firstNames)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)

	} else {

		if !isValidName {
			fmt.Println("First name or last name you entered is too short.")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign.")
		}
		if !isValidTickets {
			fmt.Println("Number of tickets you entered is invalid.")
		}

		fmt.Println("Your input data is invalid try again")
	}

	if remainingTickets == 0 {
		fmt.Println("Our conference is booked out. Come back next year.")

	}

	city := "London"

	switch city {
	case "New York":
		//execute code for booking new york conference tickets
	case "Singapore":
		//execute code for booking singapore conference tickets
	case "Mexico":
		//some cod here
	default:
		fmt.Print("no valid city selected")
	}
	wg.Wait();
}

func greetUsers() {
	fmt.Println("Welcome to", conferenceName, "booking applicaton")
	fmt.Println("We have total of", conferenceTickets, "tickets.")
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	// userName = "Himanshu"
	// userTickets = 2
	fmt.Print("Please enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your second name : ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email address : ")
	fmt.Scan(&email)

	fmt.Print("How many tickets you wanna buy : ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets int, firstName string, lastName string, email string) {

	var userData = UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		tickets:   userTickets,
	}

	remainingTickets -= userTickets
	// bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############################################")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("###############################################")
	wg.Done();
}
