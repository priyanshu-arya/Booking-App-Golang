package main

import (
	"strings"
)

func validateUserInput(fname string, lname string, email string, mobile string, bookTic uint) (bool, bool, bool, bool, bool) {
	isValidfname := len(fname) > 2
	isValidlname := len(lname) > 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidMobile := len(mobile) == 10 || len(mobile) == 12
	isValidTic := bookTic > 0 && bookTic < remainTicket
	return isValidfname, isValidlname, isValidEmail, isValidMobile, isValidTic
}
