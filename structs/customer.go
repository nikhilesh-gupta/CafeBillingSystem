package structs

import "time"

type Customer struct {
	Name      string
	Email     string
	ContactNo string
	Gender    string
	Age       int
	Time      time.Time
}
