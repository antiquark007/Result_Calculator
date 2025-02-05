package main

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"Username"`
	Password string `json:"Userpass"`
}

var User1 User

func (app *App) LoginUser(User1 User) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT Userpass 
	FROM User 
	WHERE Username = ?`
	row := app.DB.QueryRowContext(ctx, query, User1.Username)
	if err != nil {
		return "", http.StatusBadGateway, err
	}
	var Password string
	err = row.Scan(&Password)
	if err != nil {
		return "", http.StatusBadGateway, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(User1.Password), []byte(Password))
	if err != nil {
		return "", http.StatusUnauthorized, err
	}
	return User1.Username, http.StatusOK, nil
}
