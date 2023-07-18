package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	Template_Cache map[string]*template.Template
	UseCache       bool
	InProduction   bool
	Logger         *log.Logger
	Session        *scs.SessionManager
}
