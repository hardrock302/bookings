package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"bookings/pkg/render"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

var app_config config.AppConfig
var session *scs.SessionManager

func main() {
	session = scs.New()
	app_config.InProduction = false
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app_config.InProduction
	app_config.Template_Cache = render.CreateTemplateCache()
	app_config.UseCache = true
	app_config.Session = session
	port := ":8080"

	handlers.NewRepo(&app_config)
	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app_config),
	}
	err := serve.ListenAndServe()
	log.Println(fmt.Sprintf("Application running on port %s", port))
	if err != nil {
		log.Fatal(err)
	}
}
