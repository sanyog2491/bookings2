package repository

import (
	"time"

	"github.com/sanyog2491/bookings2/internal/models"
)

type DatabaseRepo interface {
	AllUsers() bool

	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDatesByRoomID(start, end time.Time) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
}
