package dto

type ReservationRequest struct {
	Restaurant    int    `json:"restaurant"`
	CustomerName  string `json:"customer_name"`
	CustomerPhone string `json:"customer_phone"`
	Table         int    `json:"table" db:"table_id"`
	Time          string `json:"time"`
}
