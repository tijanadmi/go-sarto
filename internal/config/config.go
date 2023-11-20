package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
	"github.com/tijanadmi/go-sarto/internal/models"
)

// AppConfig holds the application config

type MailServer struct{
	Host string
	PortNumber  string
	Username string
	Password string

}

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	MailChan      chan models.MailData
	MailServer MailServer
}
