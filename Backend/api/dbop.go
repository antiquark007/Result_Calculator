package main

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"userpass"`
}
type registerUser struct {
	Username string `json:"username"`
	Password string `json:"userpass"`
	Semester int    `json:"semester"`
}

var User1 User

func (app *App) LoginUser(User1 User) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	query := `SELECT userpass 
	FROM "user" 
	WHERE username = $1`
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
func (app *App) Register(User1 registerUser) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var existingUser string
	queryCheck := `SELECT username FROM "user" WHERE username = $1`
	err := app.DB.QueryRowContext(ctx, queryCheck, User1.Username).Scan(&existingUser)
	if err != nil && err != sql.ErrNoRows {

		return "", http.StatusInternalServerError, err
	}

	if existingUser != "" {

		return "Username already taken", http.StatusConflict, nil
	}

	queryInsert := `INSERT INTO "user" (username, userpass, semester) VALUES ($1, $2, $3)`
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(User1.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}
	_, err = app.DB.ExecContext(ctx, queryInsert, User1.Username, hashedPassword, User1.Semester)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return "User registered successfully", http.StatusOK, nil
}
