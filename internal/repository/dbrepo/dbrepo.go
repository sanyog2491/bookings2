package dbrepo

import (
	"database/sql"

	"github.com/sanyog2491/bookings2/internal/config"
	"github.com/sanyog2491/bookings2/internal/repository"
)

type postgressDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewpostgressRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgressDBRepo{
		App: a,
		DB:  conn,
	}
}
