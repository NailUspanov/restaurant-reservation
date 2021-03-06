package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"restaurant-reservation/internal/domain"
	"strings"
)

type TablePostgres struct {
	db *sqlx.DB
}

func NewTablePostgres(db *sqlx.DB) *TablePostgres {
	return &TablePostgres{db: db}
}

func (t *TablePostgres) GetAllNotIn(args []int) ([]domain.Table, error) {
	var tables []domain.Table

	count := len(args) - 1
	arr := make([]any, count)
	for i := range arr {
		arr[i] = i + 2
	}
	selectAvailableTablesQuery :=
		fmt.Sprintf("select * from "+TablesTable+" where id not in ($1"+strings.Repeat(",$%d", count)+")", arr...)

	s := make([]interface{}, len(args))
	for i, j := range args {
		s[i] = j
	}

	err := t.db.Select(&tables, selectAvailableTablesQuery, s...)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (t *TablePostgres) GetAllByRestaurant(restaurantId int) ([]domain.Table, error) {
	var tables []domain.Table
	getRestaurantTablesQuery := fmt.Sprintf("SELECT t.* FROM %s t WHERE t.restaurant=$1", TablesTable)
	err := t.db.Select(&tables, getRestaurantTablesQuery, restaurantId)
	if err != nil {
		return nil, err
	}

	return tables, nil
}

func (t *TablePostgres) GetAllRestaurantsCapacity() (map[int]int, error) {
	result := make(map[int]int)
	getAllQuery := fmt.Sprintf("SELECT t.restaurant, SUM(capacity) as capacity FROM %s t GROUP BY t.restaurant", TablesTable)
	rows, err := t.db.QueryxContext(context.TODO(), getAllQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var restId int
		var capacity int
		err := rows.Scan(&restId, &capacity)
		if err != nil {
			return nil, err
		}
		result[restId] = capacity
	}

	return result, nil
}
