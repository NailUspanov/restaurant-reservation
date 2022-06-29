package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CustomerPostgres struct {
	db *sqlx.DB
}

func NewCustomerPostgres(db *sqlx.DB) *CustomerPostgres {
	return &CustomerPostgres{db: db}
}

func (c *CustomerPostgres) GetCustomerIdByPhone(phone string) (int, error) {
	var customerId int

	selectCustomerQuery := fmt.Sprintf("SELECT c.id FROM %s c WHERE c.phone=$1", CustomersTable)
	row := c.db.QueryRow(selectCustomerQuery, phone)
	if err := row.Scan(&customerId); err != nil {
		return 0, err
	}

	return customerId, nil
}

func (c *CustomerPostgres) Create(name string, phone string) (int, error) {
	var customerId int

	createCustomerQuery := fmt.Sprintf("INSERT INTO %s (name, phone) VALUES ($1, $2) RETURNING id", CustomersTable)
	row := c.db.QueryRow(createCustomerQuery, name, phone)
	if err := row.Scan(&customerId); err != nil {
		return 0, err
	}

	return customerId, nil
}
