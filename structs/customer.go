package structs

type Customer struct {
	Name        string
	Email       string
	ContactNo   string
	Gender      string
	Age         int
	Time        TimeFormat
	Order       []Order
	TotalAmount int
}

type TimeFormat struct {
	Day  string
	Date string
	Time string
}
