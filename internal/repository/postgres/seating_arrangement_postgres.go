package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SeatingArrangementPostgres struct {
	db *sqlx.DB
}

func NewSeatingArrangementPostgres(db *sqlx.DB) *SeatingArrangementPostgres {
	return &SeatingArrangementPostgres{db: db}
}

func (s *SeatingArrangementPostgres) Create(tableId int, reservationId int) (int, error) {
	var id int

	createSeatingArrangementQuery := fmt.Sprintf(`INSERT INTO %s ("table", reservation) VALUES ($1, $2) RETURNING id`, SeatingArrangementTable)
	row := s.db.QueryRow(createSeatingArrangementQuery, tableId, reservationId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
