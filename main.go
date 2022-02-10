package main

import (
	"fmt"
	"log"

	"github.com/nikhilesh-gupta/CafeBillingSystem/users"
)

func init() {
	fmt.Println("-----------------------")
	fmt.Println("Welcome to XYZ Cafe !!!")
	fmt.Println("-----------------------")
}

func main() {

	// Variables
	var endUser string
	var invalidCounter int

	//Start of the Program
START:
	fmt.Print("You are Customer or Admin. Type:[C/A] >> ")
	fmt.Scanln(&endUser)

	if endUser == "C" || endUser == "c" || endUser == "customer" || endUser == "Customer" {
		users.CustomerDetails()
		users.ShowMenu()
		users.TakeOrder()

	} else if endUser == "A" || endUser == "a" || endUser == "admin" || endUser == "Admin" {

	} else {
		if invalidCounter == 3 {
			log.Fatal("\n[X]EXITING: Too many invalid inputs.")
		}
		fmt.Println("[!]WARNING: Invalid Input. Please try again!!!")
		invalidCounter++
		goto START
	}
}
