package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/pkg/models"
	"strconv"
	"strings"
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

func (r *RestaurantPostgres) GetAvailable(peopleQuantity int, time string) ([]models.AvailableRestaurantResponse, error) {

	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	var reservations []models.Reservation

	// получил все брони в указанное время
	selectReservationsQuery := fmt.Sprintf("SELECT r.id, r.restaurant, r.customer, r.table_id, lower(r.time), upper(r.time) FROM %s r WHERE time@>$1", reservationTable)
	rows, err := r.db.Query(selectReservationsQuery, time)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//err = r.db.Select(&reservations, selectReservationsQuery, time)
	for rows.Next() {
		var reservation models.Reservation
		err = rows.Scan(&reservation.Id, &reservation.Restaurant, &reservation.Customer, &reservation.Table, &reservation.Time[0], &reservation.Time[1])
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	unavailableRestaurantTables := make(map[int][]int)
	availableRestaurantTables := make(map[int][]models.Table)
	availableRestaurantResponse := make([]models.AvailableRestaurantResponse, 0, 5)

	// если брони в указанное время есть, проверяю есть ли достаточное количество мест для новой брони
	if len(reservations) > 0 {

		for _, v := range reservations {
			// инициализирую пустыми массимвами
			if unavailableRestaurantTables[v.Restaurant] == nil {
				unavailableRestaurantTables[v.Restaurant] = make([]int, 0, 3)
				availableRestaurantTables[v.Restaurant] = make([]models.Table, 0, 3)
			}
			// добавляю занятый стол в мапу
			unavailableRestaurantTables[v.Restaurant] = append(unavailableRestaurantTables[v.Restaurant], v.Table)
		}

		// заполняю мапу свободных столов
		for k, v := range unavailableRestaurantTables {
			count := len(v) - 1
			arr := make([]any, count)
			for i := range arr {
				arr[i] = i + 2
			}
			selectAvailableTablesQuery :=
				fmt.Sprintf("select * from "+tablesTable+" where id not in ($1"+strings.Repeat(",$%d", len(v)-1)+")", arr...)

			s := make([]interface{}, len(v))
			for i, j := range v {
				s[i] = j
			}
			tables := availableRestaurantTables[k]
			err := r.db.Select(&tables, selectAvailableTablesQuery, s...)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			for _, i2 := range tables {
				availableRestaurantTables[i2.Restaurant] = append(availableRestaurantTables[i2.Restaurant], i2)
			}
		}

		// из множества доступных столов вычислить общее количество доступных мест
		for k, v := range availableRestaurantTables {
			seats := 0
			for _, table := range v {
				seats += table.Capacity
			}
			if peopleQuantity > seats {
				delete(availableRestaurantTables, k)
			}
		}

		for k, v := range availableRestaurantTables {
			var restaurant models.Restaurant
			getByIdQuery := fmt.Sprintf("SELECT r.* FROM %s r WHERE r.id = $1", restaurantTable)
			err := r.db.Get(&restaurant, getByIdQuery, k)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
			availableRestaurantResponse = append(availableRestaurantResponse, models.AvailableRestaurantResponse{
				Name:            restaurant.Name,
				Location:        restaurant.Location,
				AvgWaitingTime:  restaurant.AvgWaitingTime,
				AvgBillAmount:   restaurant.AvgBillAmount,
				AvailableTables: v,
			})
		}

	} else { // если броней в указанное время нет, проверяю вместимость доступных ресторанов для компании
		getAllQuery := fmt.Sprintf("SELECT t.restaurant, SUM(capacity) as capacity FROM %s t GROUP BY t.restaurant", tablesTable)
		rows, err := r.db.QueryxContext(context.TODO(), getAllQuery)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		for rows.Next() {

			var availableRestaurant models.AvailableRestaurantResponse
			var seats int

			err = rows.Scan(
				&availableRestaurant.Name,
				&seats,
			)

			if err != nil {
				tx.Rollback()
				return nil, err
			}
			if peopleQuantity <= seats {
				var tables []models.Table
				getRestaurantTablesQuery := fmt.Sprintf("SELECT t.* FROM %s t WHERE t.restaurant=$1", tablesTable)
				err := r.db.Select(&tables, getRestaurantTablesQuery, availableRestaurant.Name)
				if err != nil {
					return nil, err
				}
				restId, err := strconv.Atoi(availableRestaurant.Name)
				if err != nil {
					return nil, err
				}
				restaurant, err := r.GetById(restId)
				if err != nil {
					return nil, err
				}

				availableRestaurant.Name = restaurant.Name
				availableRestaurant.Location = restaurant.Location
				availableRestaurant.AvgBillAmount = restaurant.AvgBillAmount
				availableRestaurant.AvgWaitingTime = restaurant.AvgWaitingTime
				availableRestaurant.AvailableTables = tables

				availableRestaurantResponse = append(availableRestaurantResponse, availableRestaurant)
			}
		}
	}

	return availableRestaurantResponse, tx.Commit()
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
