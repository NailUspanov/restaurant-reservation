package repository

import (
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/internal/domain"
	"restaurant-reservation/internal/domain/dto"
	"restaurant-reservation/internal/repository/postgres"
)

type Reservation interface {
	Create(reservation dto.ReservationRequest) (int, error)
	GetAll(customerId int) ([]domain.Reservation, error)
	GetById(reservationId int) (domain.Reservation, error)
	Delete(reservationId int) error
	GetAllByTime(time string) ([]domain.Reservation, error)
}

type Restaurant interface {
	Create(restaurant domain.Restaurant) (int, error)
	GetAll() ([]domain.Restaurant, error)
	GetById(restaurantId int) (domain.Restaurant, error)
	GetByIds(restaurantIds []int) ([]domain.Restaurant, error)
	Delete(restaurantId int) error
}

type Table interface {
	GetAllNotIn(args []int) ([]domain.Table, error)
	GetAllByRestaurant(restaurantId int) ([]domain.Table, error)
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
		Reservation:        postgres.NewReservationPostgres(db),
		Restaurant:         postgres.NewRestaurantPostgres(db),
		Table:              postgres.NewTablePostgres(db),
		Customer:           postgres.NewCustomerPostgres(db),
		SeatingArrangement: postgres.NewSeatingArrangementPostgres(db),
	}
}
