package main

import (
	"encoding/json"
	"net/http"
)

func (app *App) Home(w http.ResponseWriter, r*http.Request) {
	type response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

	resp := response{
		Message: "Hello",
		Status:  200,
	}
	json.NewEncoder(w).Encode(resp)
}