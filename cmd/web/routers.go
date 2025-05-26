package main

import (
	"net/http"

	"github.com/fahedul-islam/bookings/pkg/config"
	"github.com/fahedul-islam/bookings/pkg/handler"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig) http.Handler {
	mux:= chi.NewRouter()

	 //mux.Use(WriteToConsole)
	 mux.Use(NoSurf)
	 mux.Use(SessionLoad)
	mux.Get("/",handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)

	return mux
}