package main

import "net/http"

func (app *App) enableCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "http://localhost:5173" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept, X-CSRF-Token, Authorization")

		}

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		} else {

			h.ServeHTTP(w, r)
		}
	})
}
