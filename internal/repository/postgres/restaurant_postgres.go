package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/internal/domain"
	"strings"
)

type RestaurantPostgres struct {
	db *sqlx.DB
}

func NewRestaurantPostgres(db *sqlx.DB) *RestaurantPostgres {
	return &RestaurantPostgres{db: db}
}

func (r *RestaurantPostgres) Create(restaurant domain.Restaurant) (int, error) {
	return 0, errors.New("")
}

func (r *RestaurantPostgres) GetAll() ([]domain.Restaurant, error) {
	return nil, errors.New("")
}

func (r *RestaurantPostgres) GetById(restaurantId int) (domain.Restaurant, error) {
	var restaurant domain.Restaurant
	getByIdQuery := fmt.Sprintf("SELECT r.* FROM %s r WHERE r.id = $1", RestaurantTable)
	err := r.db.Get(&restaurant, getByIdQuery, restaurantId)

	return restaurant, err
}

func (r *RestaurantPostgres) GetByIds(restaurantIds []int) ([]domain.Restaurant, error) {
	var restaurants []domain.Restaurant

	count := len(restaurantIds) - 1
	arr := make([]any, count)
	for i := range arr {
		arr[i] = i + 2
	}
	selectRestaurantsQuery := fmt.Sprintf("select * from "+RestaurantTable+" where id in ($1"+strings.Repeat(",$%d", count)+") ORDER BY avg_waiting_time, avg_bill_amount", arr...)

	s := make([]interface{}, len(restaurantIds))
	for i, j := range restaurantIds {
		s[i] = j
	}

	err := r.db.Select(&restaurants, selectRestaurantsQuery, s...)

	return restaurants, err
}

func (r *RestaurantPostgres) Delete(restaurantId int) error {
	return errors.New("")
}
