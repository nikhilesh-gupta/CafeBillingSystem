package bills

import (
	"fmt"
	"log"
	"strconv"

	"github.com/nikhilesh-gupta/CafeBillingSystem/structs"
)

var totalAmount int

func dishDetails(dish string, cust []structs.Customer, menu structs.Menu) (name string, price int) {
	switch dish {
	case "C":
		name = menu.C.Dish
		price = menu.C.Price

	case "D":
		name = menu.D.Dish
		price = menu.D.Price

	case "T":
		name = menu.T.Dish
		price = menu.T.Price

	case "I":
		name = menu.I.Dish
		price = menu.I.Price

	case "V":
		name = menu.V.Dish
		price = menu.V.Price

	case "B":
		name = menu.B.Dish
		price = menu.B.Price

	case "P":
		name = menu.P.Dish
		price = menu.P.Price

	case "M":
		name = menu.M.Dish
		price = menu.M.Price

	case "H":
		name = menu.H.Dish
		price = menu.H.Price

	case "F":
		name = menu.F.Dish
		price = menu.F.Price

	case "J":
		name = menu.J.Dish
		price = menu.J.Price

	case "L":
		name = menu.L.Dish
		price = menu.L.Price

	case "S":
		name = menu.S.Dish
		price = menu.S.Price

	default:
		name = "nil"
		price = 0
	}

	return
}

func GenerateBill(cust []structs.Customer, menu structs.Menu) {
	fmt.Println("\n------------")
	fmt.Println("BILL")
	fmt.Println("------------")

	if cust[len(cust)-1].Gender == "M" {
		fmt.Println("Customer Name: Mr.", cust[len(cust)-1].Name)
	} else if cust[len(cust)-1].Gender == "F" {
		fmt.Println("Customer Name: Mrs.", cust[len(cust)-1].Name)
	} else {
		fmt.Println("Customer Name: ", cust[len(cust)-1].Name)
	}

	fmt.Println("Your Orders:")

	func() { //Anonymous function to calculate total amount
		for _, val := range cust[len(cust)-1].Order {
			quantity, err := strconv.Atoi(val.Quantity)
			name, price := dishDetails(string(val.Dish), cust, menu)
			fmt.Printf("\t%s x%d\n", name, quantity)
			if err != nil {
				log.Fatal("\n[!]ERROR: ", err)
			}
			totalAmount += quantity * price
		}
	}()
	fmt.Println("_______________________________")
	fmt.Println("Total Amount: Rs.", totalAmount)

	fmt.Println("\nThankyou for buying from XYZ Cafe. Please come here again ")
}
