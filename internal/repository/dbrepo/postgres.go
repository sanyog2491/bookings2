package dbrepo

import (
	"context"
	"time"

	"github.com/sanyog2491/bookings2/internal/models"
)

func (m *postgressDBRepo) AllUsers() bool {
	return true
}

func (m *postgressDBRepo) InsertReservation(res models.Reservation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `insert into reservations(first_name,last_name,email,phone,start_date,end_date,room_id,created_at,updated_at)
			values ($1,$2,$3,$4,$5,$6,$7,$8,$9)`

	_, err := m.DB.ExecContext(ctx, query,
		res.Firstname,
		res.Lastname,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil

}
