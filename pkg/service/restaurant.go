package service

import (
	"restaurant-reservation/pkg/models"
	"restaurant-reservation/pkg/repository"
)

type RestaurantService struct {
	repo repository.Restaurant
}

func NewRestaurantService(repo repository.Restaurant) *RestaurantService {
	return &RestaurantService{repo: repo}
}

func (r *RestaurantService) Create(restaurant models.Restaurant) (int, error) {
	return r.repo.Create(restaurant)
}

func (r *RestaurantService) GetAvailable(peopleQuantity int, time string) ([]models.AvailableRestaurantResponse, error) {
	return r.repo.GetAvailable(peopleQuantity, time)
}

func (r *RestaurantService) GetAll() ([]models.Restaurant, error) {
	return r.repo.GetAll()
}

func (r *RestaurantService) GetById(restaurantId int) (models.Restaurant, error) {
	return r.repo.GetById(restaurantId)
}

func (r *RestaurantService) Delete(restaurantId int) error {
	return r.repo.Delete(restaurantId)
}
