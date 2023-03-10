package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/sanyogpatel-tecblic/bookings/internal/models"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	Mailchan      chan models.MailData
}
