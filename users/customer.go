package users

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/nikhilesh-gupta/CafeBillingSystem/structs"
)

// Global Variables
var cust []structs.Customer
var menu structs.Menu
var custOrder []structs.Order

// var details []structs.Customer

func CustomerDetails(scanner *bufio.Reader) {
	//Variables
	var name, email, contactNo, gender string
	var age int
	_, err := scanner.ReadString('\n')
	errorCheck(err)

	fmt.Println("Please enter your details [* one's are mandatory]")
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Contact Number: ")
	fmt.Scanln(&contactNo)
	fmt.Print("Gender[M/F]: ")
	fmt.Scanln(&gender)
	fmt.Print("Age: ")
	fmt.Scanln(&age)

	// Taking the orders
	takeOrder()

	if len(custOrder) == 0 { //If nothing is ordered
		log.Fatal("\n[!]WARNING: \n\tSorry but you haven't ordered anything. \n\t[<NO BILL GENERATED>] \n\tPlease start the program again to order.")
	}

	//Unmarshalling
	customerJson, err := os.Open("data/customerDetails.json")
	errorCheck(err)
	defer customerJson.Close()

	customerJsonBytes, err := ioutil.ReadAll(customerJson)
	errorCheck(err)

	if len(customerJsonBytes) != 0 { //If the json file is empty then don't unmarshal it
		err = json.Unmarshal(customerJsonBytes, &cust)
		errorCheck(err)
	}

	// For Time Formatting
	timeFormat := time.Now()

	cust = append(cust, structs.Customer{
		Name:      name,
		Email:     email,
		ContactNo: contactNo,
		Gender:    gender,
		Age:       age,
		Time: structs.TimeFormat{
			Day:  timeFormat.Format("Monday"),
			Date: timeFormat.Format("01-02-2006"),
			Time: timeFormat.Format("15:04:05"),
		},
		Order: custOrder,
	})

	storeCustomerDetails()

}

func storeCustomerDetails() {
	details, err := json.MarshalIndent(cust, "", " ")
	errorCheck(err)

	//Writing into JSON file
	err = ioutil.WriteFile("data/customerDetails.json", details, 0644)
	errorCheck(err)
}

//To check any unexcepted errors in the program
func errorCheck(err error) {
	if err != nil {
		log.Fatal("\n[!]ERROR: ", err)
	}
}

func ShowMenu() {
	menuJson, err := os.Open("data/menu.json")
	errorCheck(err)
	defer menuJson.Close()

	menuJsonBytes, err := ioutil.ReadAll(menuJson)
	errorCheck(err)

	err = json.Unmarshal(menuJsonBytes, &menu)
	errorCheck(err)
}

func takeOrder() {
	var order structs.Order
	for {
		fmt.Print("Order: ")
		fmt.Scanf("%s %s", &order.Quantity, &order.Dish)
		if order.Quantity == "End" || order.Quantity == "end" || order.Quantity == "q" || order.Quantity == "quit" {
			break
		}
		custOrder = append(custOrder, order)
	}
}
