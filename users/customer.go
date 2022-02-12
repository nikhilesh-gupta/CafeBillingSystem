package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/nikhilesh-gupta/CafeBillingSystem/bills"
	"github.com/nikhilesh-gupta/CafeBillingSystem/structs"
)

// Global Variables
var cust []structs.Customer
var menu structs.Menu
var custOrder []structs.Order

// var details []structs.Customer

func CustomerDetails() {
	//Variables
	var name, firstName, lastName, email, contactNo, gender string
	var age int

	fmt.Println("Please enter your details -")
	fmt.Print("Name: ")
	fmt.Scanf("%s%s", &firstName, &lastName)
	name = firstName + " " + lastName
	fmt.Print("Email: ")
	fmt.Scanln(&email)
	fmt.Print("Contact Number: ")
	fmt.Scanln(&contactNo)

	fmt.Print("Gender[M/F]: ")
	fmt.Scanln(&gender)

	// Validating Gender
	if gender == "M" || gender == "m" {
		gender = "M"
	} else if gender == "F" || gender == "f" {
		gender = "F"
	} else {
		fmt.Println("\n[!]WARNING: Please enter a valid type in Gender. Type 'M' or 'F'")
		log.Fatal("\n[X]Error: Program flow is break, please restart the program to continue")
	}

	fmt.Print("Age: ")
	fmt.Scanln(&age)
	// Validating Age
	if age < 4 || age > 121 {
		fmt.Println("\n[!]WARNING: Please enter a valid Age range between 4yrs - 120yrs")
		log.Fatal("\n[X]Error: Program flow is break, please restart the program to continue")
	}

	defer storeCustomerDetails()

	// Show Menu
	showMenu()
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
		Order:       custOrder,
		TotalAmount: 0,
	})

	bills.GenerateBill(cust, menu)

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

func showMenu() {
	menuJson, err := os.Open("data/menu.json")
	errorCheck(err)
	defer menuJson.Close()

	menuJsonBytes, err := ioutil.ReadAll(menuJson)
	errorCheck(err)

	err = json.Unmarshal(menuJsonBytes, &menu)
	errorCheck(err)

	fmt.Println("\n-----------")
	fmt.Println("MENU ")
	fmt.Println("-----------")
	// fmt.Println("Name \t\t\t Price")
	fmt.Printf("%v - Rs.%v\n", menu.C.Dish, menu.C.Price)
	fmt.Printf("%v - Rs.%v\n", menu.D.Dish, menu.D.Price)
	fmt.Printf("%v - Rs.%v\n", menu.T.Dish, menu.T.Price)
	fmt.Printf("%v - Rs.%v\n", menu.I.Dish, menu.I.Price)
	fmt.Printf("%v - Rs.%v\n", menu.V.Dish, menu.V.Price)
	fmt.Printf("%v - Rs.%v\n", menu.B.Dish, menu.B.Price)
	fmt.Printf("%v - Rs.%v\n", menu.P.Dish, menu.P.Price)
	fmt.Printf("%v - Rs.%v\n", menu.M.Dish, menu.M.Price)
	fmt.Printf("%v - Rs.%v\n", menu.H.Dish, menu.H.Price)
	fmt.Printf("%v - Rs.%v\n", menu.F.Dish, menu.F.Price)
	fmt.Printf("%v - Rs.%v\n", menu.J.Dish, menu.J.Price)
	fmt.Printf("%v - Rs.%v\n", menu.L.Dish, menu.L.Price)
	fmt.Printf("%v - Rs.%v\n", menu.S.Dish, menu.S.Price)
}

func takeOrder() {
	var order structs.Order
	for {
		fmt.Print("\nOrder: ")
		fmt.Scanf("%s %s", &order.Quantity, &order.Dish)
		if order.Quantity == "End" || order.Quantity == "end" || order.Quantity == "q" || order.Quantity == "quit" {
			break
		}
		custOrder = append(custOrder, order)
	}
}
