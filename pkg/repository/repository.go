package repository

import (
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/pkg/models"
)

type Reservation interface {
	Create(reservation models.ReservationRequest) (int, error)
	GetAll(customerId int) ([]models.Reservation, error)
	GetById(reservationId int) (models.Reservation, error)
	Delete(reservationId int) error
}

type Restaurant interface {
	Create(restaurant models.Restaurant) (int, error)
	GetAvailable(peopleQuantity int, time string) ([]models.AvailableRestaurantResponse, error)
	GetAll() ([]models.Restaurant, error)
	GetById(restaurantId int) (models.Restaurant, error)
	Delete(restaurantId int) error
}

type Repository struct {
	Reservation
	Restaurant
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reservation: NewReservationPostgres(db),
		Restaurant:  NewRestaurantPostgres(db),
	}
}
