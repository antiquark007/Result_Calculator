package main

import (
	"log"
	"net/http"
)

type App struct {
	Domain string
}

var err error

func main() {
	var app App
	app.Domain = "example.com"

	err = http.ListenAndServe(":6969", app.router())
	if err != nil {
		log.Fatal(err)
	}
}
