package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

func NoSurf(next http.Handler) http.Handler {
	csrf_handler := nosurf.New(next)
	csrf_handler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app_config.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrf_handler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
