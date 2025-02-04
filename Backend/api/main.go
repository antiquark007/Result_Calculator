package main

import (
	"database/sql"
	"log"
	"net/http"
)

type App struct {
	Domain string
	DB     *sql.DB
}

var err error

func main() {
	var app App
	app.Domain = "example.com"
	app.DB,err = app.conntodb()
	if err!=nil{
		log.Fatal(err)
	}
	err = http.ListenAndServe(":6969", app.router())
	if err != nil {
		log.Fatal(err)
	}
}
