package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/pkg/models"
)

type ReservationPostgres struct {
	db *sqlx.DB
}

func NewReservationPostgres(db *sqlx.DB) *ReservationPostgres {
	return &ReservationPostgres{db: db}
}

func (r *ReservationPostgres) Create(reservation models.ReservationRequest) (int, error) {
	//начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var customerId int

	//проверяю, есть ли в бд customer с таким номером телефона: если нет, добавляю
	selectCustomerQuery := fmt.Sprintf("SELECT c.id FROM %s c WHERE c.phone=$1", customersTable)
	row := tx.QueryRow(selectCustomerQuery, reservation.CustomerPhone)
	if err := row.Scan(&customerId); err != nil {
		createCustomerQuery := fmt.Sprintf("INSERT INTO %s (name, phone) VALUES ($1, $2) RETURNING id", customersTable)
		row = tx.QueryRow(createCustomerQuery, reservation.CustomerName, reservation.CustomerPhone)
		if err := row.Scan(&customerId); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	//добавляю запись в таблицу с бронями
	var reservationId int
	createReservationQuery := fmt.Sprintf("INSERT INTO %s (restaurant, customer, table_id, time) VALUES ($1, $2, $3, $4) RETURNING id", reservationTable)
	row = tx.QueryRow(createReservationQuery, reservation.Restaurant, customerId, reservation.Table, reservation.Time)
	if err := row.Scan(&reservationId); err != nil {
		tx.Rollback()
		return 0, err
	}

	//добавляю запись в таблицу seating_arrangements
	createSeatingArrangementQuery := fmt.Sprintf(`INSERT INTO %s ("table", reservation) VALUES ($1, $2) RETURNING id`, seatingArrangementTable)
	_, err = tx.Exec(createSeatingArrangementQuery, reservation.Table, reservationId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	//завершаю транзацкию
	return reservationId, tx.Commit()
}

func (r *ReservationPostgres) GetAll(customerId int) ([]models.Reservation, error) {
	var reservations []models.Reservation
	getAllQuery := fmt.Sprintf("SELECT rsrv.* FROM %s rsrv WHERE rsrv.customer=$1", reservationTable)
	err := r.db.Select(&reservations, getAllQuery, customerId)
	if err != nil {
		return nil, err
	}

	return reservations, nil
}

func (r *ReservationPostgres) GetById(reservationId int) (models.Reservation, error) {
	return models.Reservation{}, errors.New("getById not implemented")
}

func (r *ReservationPostgres) Delete(reservationId int) error {
	return errors.New("delete not implemented")
}

func (r *ReservationPostgres) GetAllByTime(time string) ([]models.Reservation, error) {
	var reservations []models.Reservation

	// получил все брони, время которых пересекается с указанным временем
	selectReservationsQuery := fmt.Sprintf("SELECT r.id, r.restaurant, r.customer, r.table_id, lower(r.time), upper(r.time) FROM %s r WHERE time&&$1", reservationTable)
	rows, err := r.db.Query(selectReservationsQuery, time)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var reservation models.Reservation
		err = rows.Scan(&reservation.Id, &reservation.Restaurant, &reservation.Customer, &reservation.Table, &reservation.Time[0], &reservation.Time[1])
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
