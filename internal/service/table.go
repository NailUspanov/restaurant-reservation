package service

import (
	"restaurant-reservation/internal/domain"
	"restaurant-reservation/internal/repository"
)

type TableService struct {
	repos repository.Table
}

func (t *TableService) GetAllRestaurantsCapacity() (map[int]int, error) {
	return t.repos.GetAllRestaurantsCapacity()
}

func NewTableService(repos repository.Table) *TableService {
	return &TableService{repos: repos}
}

func (t *TableService) GetAllNotIn(args []int) ([]domain.Table, error) {
	return t.repos.GetAllNotIn(args)
}

func (t *TableService) GetAllByRestaurant(restaurantId int) ([]domain.Table, error) {
	return t.repos.GetAllByRestaurant(restaurantId)
}
