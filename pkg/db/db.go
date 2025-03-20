package db

import "database/sql"

type Connection struct {
	*sql.DB
}

func NewConnection(DSN string) (*Connection, error) {
	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}
	return &Connection{db}, nil
}
