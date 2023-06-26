package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	//"gomail"
	//"gopkg.in/gomail.v2"
	//"sync"
	//"strconv"
	//"strings"
)

var conferenceName = "Python Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

//var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsletter bool
}

//var wg = sync.WaitGroup{}

func main() {

	greetings()

	for {

		firstName, lastName, email, userTickets := userInput()

		isValidEmail, isValidName, isValidTicketNumber := helper.ValidateUser(firstName, lastName, email, userTickets, remainingTickets)

		if isValidEmail && isValidName && isValidTicketNumber {

			//remainingTickets = remainingTickets - userTickets
			//bookings[0] = firstName + " " + lastName
			//bookings = append(bookings, firstName+" "+lastName)
			/*fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("the first value: %v\n", bookings[0])
			fmt.Printf("Slice type: %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))
			*/

			bookTickets(remainingTickets, userTickets, firstName, lastName, email, conferenceName)
			//wg.Add(1)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("the first name of bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Print("Sorry all the tickets are booked. Come some another time")
				break
			}

		} else {

			if !isValidEmail {
				fmt.Println("Your email does not contain @ sign.")
			}
			if !isValidName {
				fmt.Println("Your name is not valid please enter a valid name.")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is not valid.")
			}

		}
		//wg.Wait()
		//gomail.Send(email, userTickets)

	}

	/*city := "London"

	switch city {
	case "New York":
		//enter code here

	case "Mexico":
		//enter code here

	case "London", "Berlin":
		//enter code here

	case "Washington":
		//enter code here

	default:
		fmt.Print("No valid city selected.")

	}*/
}

func greetings() {
	fmt.Printf("WELCOME TO OUR %v BOOKING APPLICATION\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func printFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		//names[0]

		//firstNames = append(firstNames, booking["firstName"])
		firstNames = append(firstNames, booking.firstName)

	}
	return firstNames

}

func userInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets

	//for structure
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	//create a map for user
	/*var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)*/

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve your confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var tickets = fmt.Sprintf("%v tickets for %v %v \n", userTickets, firstName, lastName)
	fmt.Println("###########")
	fmt.Printf("Sending tickets to:\n %v to email address at %v\n", tickets, email)
	fmt.Println("###########")
	//wg.Done()
}

/*
concurrency in go:
used to manage multiple tasks at the same time.
*/
