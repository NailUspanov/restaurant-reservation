package service

import (
	"restaurant-reservation/pkg/models"
	"restaurant-reservation/pkg/repository"
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

func (t *TableService) GetAllNotIn(args []int) ([]models.Table, error) {
	return t.repos.GetAllNotIn(args)
}

func (t *TableService) GetAllByRestaurant(restaurantId int) ([]models.Table, error) {
	return t.repos.GetAllByRestaurant(restaurantId)
}
