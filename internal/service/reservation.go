package service

import (
	"restaurant-reservation/internal/domain"
	"restaurant-reservation/internal/domain/dto"
	"restaurant-reservation/internal/repository"
)

type ReservationService struct {
	repo repository.Reservation
}

func (r *ReservationService) GetAllByTime(time string) ([]domain.Reservation, error) {
	return r.repo.GetAllByTime(time)
}

func NewReservationService(repo repository.Reservation) *ReservationService {
	return &ReservationService{repo: repo}
}

func (r *ReservationService) Create(reservation dto.ReservationRequest) (int, error) {
	return r.repo.Create(reservation)
}

func (r *ReservationService) GetAll(customerId int) ([]domain.Reservation, error) {
	return r.repo.GetAll(customerId)
}

func (r *ReservationService) GetById(reservationId int) (domain.Reservation, error) {
	return r.repo.GetById(reservationId)
}

func (r *ReservationService) Delete(reservationId int) error {
	return r.repo.Delete(reservationId)
}
