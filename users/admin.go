package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/nikhilesh-gupta/CafeBillingSystem/bills"
	"github.com/nikhilesh-gupta/CafeBillingSystem/structs"
)

// Global Variables
var details []structs.Customer

func totalSalesOfTheDay() {
	// Variables
	var totalSales int
	var check string

	timeFormat := time.Now()
	check = timeFormat.Format("01-02-2006")
	for index := range details {
		if check == details[index].Time.Date {
			totalSales += details[index].TotalAmount
		}
	}

	fmt.Printf("\nTotal Sales: %v\n", totalSales)

	adminOptions()
}

func specificDaySales() {
	var totalSales int
	var check string
	fmt.Print("Enter the date[format -> MM-DD-YY]: ")
	fmt.Scanln(&check)
	for index := range details {
		if check == details[index].Time.Date {
			totalSales += details[index].TotalAmount
		}
	}

	fmt.Printf("\nTotal Sales: %v\n", totalSales)

	adminOptions()
}

func totalSales() {
	// Variables
	var totalSales int

	for index := range details {
		totalSales += details[index].TotalAmount
	}

	fmt.Printf("\nTotal Sales: %v\n", totalSales)

	adminOptions()
}

func printCustomerDetails(index int) {
	fmt.Println("\t Name: ", details[index].Name)
	fmt.Println("\t Email: ", details[index].Email)
	fmt.Println("\t Contact: ", details[index].ContactNo)
	fmt.Println("\t Email: ", details[index].Email)
	fmt.Println("\t Gender: ", details[index].Email)
	fmt.Println("\t Age: ", details[index].Email)
	fmt.Printf("\t Day: %v, Date: %v, Time: %v\n", details[index].Time.Day, details[index].Time.Date, details[index].Time.Time)
	fmt.Print("\t Orders: \n")
	for _, val := range details[index].Order {
		quantity, err := strconv.Atoi(val.Quantity)
		name, _ := bills.DishDetails(string(val.Dish), details, menu)
		fmt.Printf("\t\t%s x%d\n", name, quantity)
		errorCheck(err)
	}

	fmt.Printf("\t TotalAmount: Rs.%v\n", details[index].TotalAmount)
	fmt.Print("\n\n")
}

func customerDetailsOfTheDay() {
	// Variables
	var check string

	timeFormat := time.Now()
	check = timeFormat.Format("01-02-2006")

	fmt.Println("\nToday's Customer Details")
	fmt.Print("________________________\n\n")

	for index := range details {
		if check == details[index].Time.Date {
			printCustomerDetails(index)
		}
	}
	adminOptions()
}

func specificDayCustomerDetails() {
	// Variables
	var check string

	fmt.Print("Enter the date[format -> MM-DD-YY]: ")
	fmt.Scanln(&check)

	fmt.Println("\nCustomer Details of ", check)
	fmt.Print("________________________\n\n")

	for index := range details {
		if check == details[index].Time.Date {
			printCustomerDetails(index)
		}
	}
	adminOptions()
}

func totalCustomerDetails() {
	fmt.Println("\nAll Customer Details")
	fmt.Print("________________________\n\n")

	for index := range details {
		printCustomerDetails(index)
	}

	adminOptions()

}

func showOptions() {
	fmt.Println("\n___________________________")
	fmt.Println("Choose from the options")
	fmt.Println("----------------------------")
	fmt.Println("\n's' - print Today's Total Sales")
	fmt.Println("'S' - print Overall Sales")
	fmt.Println("'f' - print Full-Details of Today's Customers")
	fmt.Println("'F' - print Full-Details of all Customers")
	fmt.Println("'c' - prints a Specific day Sales")
	fmt.Println("'C' - prints a Specific day Customer's details")

	fmt.Println("'h' - help/show commands")
	fmt.Println("'q' - to quit")

	adminOptions()
}

func adminOptions() {
	var option string

	fmt.Print("\n>> ")
	fmt.Scanln(&option)

	//Unmarshalling
	performUnmarshalling()

	if len(details) == 0 {
		fmt.Println("[!]WARNING: No data available")
	}

	switch option {
	case "s":
		totalSalesOfTheDay()
	case "S":
		totalSales()
	case "f":
		customerDetailsOfTheDay()
	case "F":
		totalCustomerDetails()
	case "h":
		showOptions()
	case "c":
		specificDaySales()
	case "C":
		specificDayCustomerDetails()
	case "q":
		fmt.Print("\n\n")
		fmt.Println("[Admin Panel Closed]")
		fmt.Println("________________________")
	}
}

func Admin() {
	menuJson, err := os.Open("data/menu.json")
	errorCheck(err)
	defer menuJson.Close()

	menuJsonBytes, err := ioutil.ReadAll(menuJson)
	errorCheck(err)

	err = json.Unmarshal(menuJsonBytes, &menu)
	errorCheck(err)

	adminPasswordValidator()
	showOptions()
}

func adminPasswordValidator() {
	var passwd string
	var passwdFailCounter int
	fmt.Println("\nPlease enter your Password[default is 'admin'] ")
askPasswd:
	fmt.Print(">> ")
	fmt.Scanln(&passwd)
	if passwd != "admin" {
		if passwdFailCounter > 2 {
			log.Fatal("\n[X]ERROR: Too many invalid attempts. Exiting the program")
		}
		fmt.Println("\n[!]WARNING: Entered Password doesn't match!! \nPlease re-enter the password\t| Attempts left: ", 2-passwdFailCounter)
		passwdFailCounter++
		goto askPasswd

	}
	fmt.Println("\nWelcome back Admin !!")

}

func performUnmarshalling() {
	customerJson, err := os.Open("data/customerDetails.json")
	errorCheck(err)
	defer customerJson.Close()
	customerJsonBytes, err := ioutil.ReadAll(customerJson)
	errorCheck(err)
	err = json.Unmarshal(customerJsonBytes, &details)
	errorCheck(err)
}
