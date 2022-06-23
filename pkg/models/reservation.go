package models

import "time"

type Reservation struct {
	Id         int          `json:"id" db:"id"`
	Restaurant int          `json:"restaurant" db:"restaurant"`
	Customer   int          `json:"customer" db:"customer"`
	Table      int          `json:"table" db:"table_id"`
	Time       [2]time.Time `json:"time" db:"time"`
}

type ReservationRequest struct {
	Restaurant    int    `json:"restaurant"`
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	Table         int    `json:"table" db:"table_id"`
	Time          string `json:"time"`
}

type Table struct {
	Id         int `json:"id" db:"id"`
	Restaurant int `json:"restaurant" db:"restaurant"`
	Capacity   int `json:"capacity" db:"capacity"`
}

type SeatingArrangement struct {
	Id          int `json:"id" db:"id"`
	Table       int `json:"table" db:"table"`
	Reservation int `json:"reservation" db:"reservation"`
}

type Customer struct {
	Id    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Phone string `json:"phone" db:"phone"`
}

type Restaurant struct {
	Id             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Location       string `json:"location" db:"location"`
	AvgWaitingTime int    `json:"avg_waiting_time" db:"avg_waiting_time"`
	AvgBillAmount  int    `json:"avg_bill_amount" db:"avg_bill_amount"`
}

type AvailableRestaurantResponse struct {
	Name            string  `json:"name" db:"name"`
	Location        string  `json:"location" db:"location"`
	AvgWaitingTime  int     `json:"avg_waiting_time" db:"avg_waiting_time"`
	AvgBillAmount   int     `json:"avg_bill_amount" db:"avg_bill_amount"`
	AvailableTables []Table `json:"available_tables"`
}
