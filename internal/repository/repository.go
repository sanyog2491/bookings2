package repository

import "github.com/sanyog2491/bookings2/internal/models"

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) error
}
