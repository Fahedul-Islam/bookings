package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// protection to all the POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly :true,
		Path :"/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad load and saves the session on every request
func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}