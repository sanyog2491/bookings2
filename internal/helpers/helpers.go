package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/sanyog2491/bookings2/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func Clienterror(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with the status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
