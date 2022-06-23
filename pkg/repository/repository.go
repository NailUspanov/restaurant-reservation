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
	GetAllByTime(time string) ([]models.Reservation, error)
}

type Restaurant interface {
	Create(restaurant models.Restaurant) (int, error)
	GetAll() ([]models.Restaurant, error)
	GetById(restaurantId int) (models.Restaurant, error)
	Delete(restaurantId int) error
}

type Table interface {
	GetAllNotIn(args []int) ([]models.Table, error)
	GetAllByRestaurant(restaurantId int) ([]models.Table, error)
	GetAllRestaurantsCapacity() (map[int]int, error) // key - restaurant id, value - people capacity
}

type Customer interface {
	GetCustomerIdByPhone(phone string) (int, error)
	Create(name string, phone string) (int, error)
}

type SeatingArrangement interface {
	Create(tableId int, reservationId int) (int, error)
}

type Repository struct {
	Reservation
	Restaurant
	Table
	Customer
	SeatingArrangement
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Reservation:        NewReservationPostgres(db),
		Restaurant:         NewRestaurantPostgres(db),
		Table:              NewTablePostgres(db),
		Customer:           NewCustomerPostgres(db),
		SeatingArrangement: NewSeatingArrangementPostgres(db),
	}
}
