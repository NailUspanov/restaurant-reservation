package service

import "restaurant-reservation/internal/repository"

type Service struct {
	repository.Reservation
	RestaurantService
	repository.SeatingArrangement
	repository.Table
	repository.Customer
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reservation:        NewReservationService(repos.Reservation),
		RestaurantService:  *NewRestaurantService(*repos),
		SeatingArrangement: NewSeatingArrangementService(repos.SeatingArrangement),
		Table:              NewTableService(repos.Table),
		Customer:           NewCustomerService(repos.Customer),
	}
}
