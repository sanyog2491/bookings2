package helpers

import "github.com/sanyog2491/bookings2/internal/config"

var app *config.AppConfig

func Newhelpers(a *config.AppConfig) {
	app = a
}
