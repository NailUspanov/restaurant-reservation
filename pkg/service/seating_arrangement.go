package service

import "restaurant-reservation/pkg/repository"

type SeatingArrangementService struct {
	repo repository.SeatingArrangement
}

func NewSeatingArrangementService(repo repository.SeatingArrangement) *SeatingArrangementService {
	return &SeatingArrangementService{repo: repo}
}

func (s *SeatingArrangementService) Create(tableId int, reservationId int) (int, error) {
	return s.repo.Create(tableId, reservationId)
}
