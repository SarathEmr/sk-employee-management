package data

import "time"

type Employee struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Age         uint8     `json:"age"`
	Address     string    `json:"address"`
	Designation string    `json:"designation"`
	JoiningDate time.Time `json:"joining_date"`
}
