package service

import "restaurant-reservation/pkg/repository"

type Service struct {
	repository.Reservation
	repository.Restaurant
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Reservation: NewReservationService(repos.Reservation),
		Restaurant:  NewRestaurantService(repos.Restaurant),
	}
}
