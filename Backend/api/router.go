package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *App) router() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(app.enableCORS)
	mux.Get("/Home",app.Home)
	mux.Post("/Login",app.Login)
	mux.Post("/Register",app.RegisterUser)
	mux.Get("/Logout",app.Logout)
	mux.Get("/RefreshToken",app.RefreshHandler)
	return mux
}