package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pushinist/events-notification-bot/config"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(dbCfg config.DBConfig) (*Storage, error) {
	DSN := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)

	db, err := sql.Open("postgres", DSN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	_, err = db.Exec(createEventsTable)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createParticipantsTable)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createAdminsTable)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(createAdminEventTable)
	if err != nil {
		return nil, err
	}

	return &Storage{
		db: db,
	}, nil
}
