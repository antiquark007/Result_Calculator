package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func (app *App) conntodb() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbuser := os.Getenv("POSTGRES_USER")
	dbpass := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	dbhost := "localhost"
	dbport := 5433

	connstr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbuser, dbpass, dbhost, dbport, dbname)
    fmt.Print(connstr)
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the database")
	}
	
	return db, err
}
