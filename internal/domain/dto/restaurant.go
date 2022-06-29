package dto

import "restaurant-reservation/internal/domain"

type AvailableRestaurantResponse struct {
	Name            string         `json:"name" db:"name"`
	Location        string         `json:"location" db:"location"`
	AvgWaitingTime  int            `json:"avg_waiting_time" db:"avg_waiting_time"`
	AvgBillAmount   int            `json:"avg_bill_amount" db:"avg_bill_amount"`
	AvailableTables []domain.Table `json:"available_tables"`
}
