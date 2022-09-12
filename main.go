package main

import (
	"fmt"
)

type user struct {
	fname   string
	lname   string
	email   string
	mobile  string
	bookTic uint
}

// global vaariable
var remainTicket uint = 100
var userDetail []user
var adminPass string = "admin"

// main Menu
func main() {
	fmt.Println("-------------------- Ticket Management System --------------------")
	name := welcomeUser()
	menu(name)
}

// Starting Logic
func menu(name string) {
	fmt.Println("\n Please Choose from Options: \n ")
	fmt.Println("1 - Book Ticket")
	fmt.Println("2 - Check Remaining Ticket")
	fmt.Println("3 - Check GO Conference Details")
	fmt.Println("4 - Admin Menu")
	fmt.Println("5 - Exit \n ")
	fmt.Print("Enter Your Option: ")
	var option int
	fmt.Scan(&option)

	switch option {
	case 1:
		bookTicket(name)
		break
	case 2:
		remainingTicket()
		break
	case 3:
		conDetails()
		break
	case 4:
		admin(name)
		break
	case 5:
		break
	default:
		fmt.Println("Please Choose from Appropriate Options")
	}

	main() // Send back to Menu Again
}

// Ticket Booking Logic
func bookTicket(fname string) {
	var newUser user
	newUser.fname = fname
	fmt.Printf("%v, Please Enter your last name \n", fname)
	fmt.Scan(&newUser.lname)
	fmt.Printf("%v, Please Enter your email \n", fname)
	fmt.Scan(&newUser.email)
	fmt.Println("Please Enter your mobile number")
	fmt.Scan(&newUser.mobile)
	fmt.Printf("\n Remaining Ticket is %v, Please enter your desired number of Ticket in between Remaining Ticket: \n", remainTicket)
	fmt.Scan(&newUser.bookTic)
	fmt.Println("Press Y for Submit")
	var op string
	fmt.Scan(&op)
	if op == "Y" || op == "y" {
		isValidfname, isValidlname, isValidEmail, isValidMobile, isValidTic := validateUserInput(newUser.fname, newUser.lname, newUser.email, newUser.mobile, newUser.bookTic)
		if isValidfname && isValidlname && isValidEmail && isValidMobile && isValidTic {
			fmt.Printf("\nCongratulations %v, Please check your %v for Confirmation \n", fname, newUser.email)
			remainTicket -= newUser.bookTic
			userDetail = append(userDetail, newUser)
		} else {
			if !(isValidfname && isValidlname) {
				fmt.Println("Name is not Valid, Please try Again")
			}
			if !(isValidEmail) {
				fmt.Println("Email is not Valid, Please try Again")
			}
			if !(isValidMobile) {
				fmt.Println("Mobile Number is not Valid, Please try Again")
			}
			if !(isValidTic) {
				fmt.Println("Sorry, We don't have that much of seat vacant Please Try Again !!")
			}
		}
	} else {
		fmt.Println("Going to main Menu")
	}

}

// Show Remaining Ticket
func remainingTicket() {
	fmt.Printf("\nGo Conference Remaining Ticket %v \n", remainTicket)
}

// Show Conference Details
func conDetails() {
	fmt.Println("Go conference Details")
	fmt.Println("Go conference is Organized at Bangalore, India on September 15, 2022.")
	fmt.Println("Main Goal of this conference to train new minds for Go Language")
	fmt.Println("PLease mail us at goConference@gocon.com for more details.")
	fmt.Println("Thankyou for asking. Have a nice day.")
}

// admin
func admin(name string) {
	var pass string
	fmt.Printf("Hey %v, Please Enter Admin Password to Continue", name)
	fmt.Scan(&pass)

	if adminPass == pass {
		adminMenu(name)
	} else {
		fmt.Println("")
		return
	}
}

// admin
func adminMenu(name string) {
	var option uint
	fmt.Printf("Hey %v, Please Choose Appropriate option\n", name)
	fmt.Println("1 - Show Complete Users Details")
	fmt.Println("2 - Show Users Name")
	fmt.Println("3 - Show Users Email")
	fmt.Println("4 - Show Users Mobile")
	fmt.Println("5 - Check Remaining Ticket")
	fmt.Println("6 - Exit \n ")
	fmt.Scan(&option)

	switch option {
	case 1:
		userDetails()
		break
	case 2:
		userName()
		break
	case 3:
		userEmail()
		break
	case 4:
		userMobile()
		break
	case 5:
		remainingTicket()
		break
	case 6:
		break
	default:
		fmt.Println("Please Choose from Appropriate Options")
		adminMenu(name)
	}
	fmt.Println("For admin menu Press Y")
	var op string
	fmt.Scan(&op)
	if op == "Y" || op == "y" {
		adminMenu(name)
	}
	main() // Send back to Menu Again
}

// User Details
func userDetails() {
	fmt.Println("Name     Email     Mobile     Tickets")
	for _, user := range userDetail {
		fmt.Printf("%v   %v    %v    %v\n", user.fname+" "+user.lname, user.email, user.mobile, user.bookTic)
	}
}

// User Name List
func userName() {
	fmt.Println("Name")
	for _, user := range userDetail {
		fmt.Println(user.fname + " " + user.lname)
	}
}

// User Email List
func userEmail() {
	fmt.Println("Email")
	for _, user := range userDetail {
		fmt.Printf("%v\n", user.email)
	}
}

// User Mobile List
func userMobile() {
	fmt.Println("Mobile Number")
	for _, user := range userDetail {
		fmt.Printf("%v\n", user.mobile)
	}
}

// Greet Guest
func welcomeUser() string {
	var name string
	fmt.Println("Hello, Please Enter your name")
	fmt.Scan(&name)
	fmt.Printf("\nHello, %v Welcome to Go Conference \n", name)
	return name
}
