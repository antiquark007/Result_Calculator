package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (app *App) Home(w http.ResponseWriter, r *http.Request) {
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

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type Response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
		Error   error  `json:"error"`
	}
	message, status, err := app.LoginUser(user)

	response := Response{
		Message: message,
		Status:  status,
		Error:   err,
	}
	accesstoken, err := GenerateAccessToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print(accesstoken)
	refreshtoken,err:=GenerateRefreshToken(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
   fmt.Print(refreshtoken)
	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   accesstoken,
		Expires: time.Now().Add(accesstokenexp),
		Secure:  false,
		HttpOnly: true,
	})
	fmt.Print("cookie genreated")
	http.SetCookie(w,&http.Cookie{
		Name:    "refresh_token",
		Value:   refreshtoken,
		Expires: time.Now().Add(refreshtokexp),
		Secure:  false,
		HttpOnly: true,

	})

	err=json.NewEncoder(w).Encode(response)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
func (app *App)RefreshHandler(w http.ResponseWriter,r*http.Request){
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	claims,err:=validatetokens(cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	newRefreshtoken,err:=GenerateRefreshToken(claims.Username)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.SetCookie(w,&http.Cookie{
		Name:    "refresh_token",
		Value:   newRefreshtoken,
		Expires: time.Now().Add(refreshtokexp),
		Secure:  false,
		HttpOnly: true,

	})
}

func (app *App)Logout (w http.ResponseWriter,r *http.Request){
	http.SetCookie(w,GetExpiredRefershCookie())
	w.WriteHeader(http.StatusAccepted)

}