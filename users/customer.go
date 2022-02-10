package users

import (
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

// var details []structs.Customer

func CustomerDetails() {
	//Variables
	var name, email, contactNo, gender string
	var age int

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

	//Unmarshalling
	customerJson, err := os.Open("data/customerDetails.json")
	errorCheck(err)
	defer customerJson.Close()

	customerJsonBytes, err := ioutil.ReadAll(customerJson)
	errorCheck(err)

	err = json.Unmarshal(customerJsonBytes, &cust)
	errorCheck(err)

	cust = append(cust, structs.Customer{
		Name:      name,
		Email:     email,
		ContactNo: contactNo,
		Gender:    gender,
		Age:       age,
		Time:      time.Now(),
	})

	storeCustomerDetails()

}

func storeCustomerDetails() {
	details, err := json.Marshal(cust)
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

	// fmt.Println(structs.Menu.C)
}

func TakeOrder() {
	var order string
	for {
		fmt.Print("Order: ")
		fmt.Scanln(&order)
		if order == "End" || order == "end" || order == "q" || order == "quit" {
			break
		}
	}
}
