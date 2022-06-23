package service

import (
	"restaurant-reservation/pkg/models"
	"restaurant-reservation/pkg/repository"
)

type ReservationService struct {
	repo repository.Reservation
}

func (r *ReservationService) GetAllByTime(time string) ([]models.Reservation, error) {
	return r.repo.GetAllByTime(time)
}

func NewReservationService(repo repository.Reservation) *ReservationService {
	return &ReservationService{repo: repo}
}

func (r *ReservationService) Create(reservation models.ReservationRequest) (int, error) {
	return r.repo.Create(reservation)
}

func (r *ReservationService) GetAll(customerId int) ([]models.Reservation, error) {
	return r.repo.GetAll(customerId)
}

func (r *ReservationService) GetById(reservationId int) (models.Reservation, error) {
	return r.repo.GetById(reservationId)
}

func (r *ReservationService) Delete(reservationId int) error {
	return r.repo.Delete(reservationId)
}
