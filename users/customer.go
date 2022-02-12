package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nikhilesh-gupta/CafeBillingSystem/bills"
	"github.com/nikhilesh-gupta/CafeBillingSystem/structs"
)

// Global Variables
var cust []structs.Customer
var menu structs.Menu
var custOrder []structs.Order

func Fatal() {
	log.Fatal("\n[X]Error: Program flow is break, please restart the program to continue")
}

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

	// Validate Email
	if !strings.Contains(email, "@") {
		fmt.Println("\n[!]WARNING: Please enter a valid Email ID. It should contain '@'")
		Fatal()
	}

	fmt.Print("Contact Number: ")
	fmt.Scanln(&contactNo)

	// Validate Contact Number
	if _, err := strconv.Atoi(contactNo); err != nil {
		fmt.Println("\n[!]WARNING: Please enter a valid Contact number. It should only contains digits")
		Fatal()
	}

	fmt.Print("Gender[M/F]: ")
	fmt.Scanln(&gender)

	// Validating Gender
	if gender == "M" || gender == "m" {
		gender = "M"
	} else if gender == "F" || gender == "f" {
		gender = "F"
	} else {
		fmt.Println("\n[!]WARNING: Please enter a valid type in Gender. Type 'M' or 'F'")
		Fatal()
	}

	fmt.Print("Age: ")
	fmt.Scanln(&age)
	// Validating Age
	if age < 5 || age > 95 {
		fmt.Println("\n[!]WARNING: Please enter a valid Age range between 5yrs - 95yrs")
		Fatal()
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

	fmt.Print("\n\n")
	fmt.Println("[Enter the quantity then first letter of the dish with a space between them.\nLike: '4 H', '3 b']")
}

func takeOrder() {
	var order structs.Order
	for {
		fmt.Print("\nOrder: ")
		fmt.Scanf("%s %s", &order.Quantity, &order.Dish)
		if order.Quantity == "End" || order.Quantity == "end" || order.Quantity == "q" || order.Quantity == "quit" {
			break
		}
		if len(order.Quantity) > 1 {
			fmt.Println("\n[!]WARNING: Enter your order in the valid format as specified")
			log.Fatal("\n[X]Error: Program flow is break, please restart the program to continue")
		}
		if order.Dish == strings.ToLower(order.Dish) {
			order.Dish = strings.ToUpper(order.Dish)
		}
		custOrder = append(custOrder, order)
	}
}
