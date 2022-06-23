package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/pkg/models"
)

type RestaurantPostgres struct {
	db *sqlx.DB
}

func NewRestaurantPostgres(db *sqlx.DB) *RestaurantPostgres {
	return &RestaurantPostgres{db: db}
}

func (r *RestaurantPostgres) Create(restaurant models.Restaurant) (int, error) {
	return 0, errors.New("")
}

func (r *RestaurantPostgres) GetAll() ([]models.Restaurant, error) {
	return nil, errors.New("")
}

func (r *RestaurantPostgres) GetById(restaurantId int) (models.Restaurant, error) {
	var restaurant models.Restaurant
	getByIdQuery := fmt.Sprintf("SELECT r.* FROM %s r WHERE r.id = $1", restaurantTable)
	err := r.db.Get(&restaurant, getByIdQuery, restaurantId)

	return restaurant, err
}

func (r *RestaurantPostgres) Delete(restaurantId int) error {
	return errors.New("")

}
