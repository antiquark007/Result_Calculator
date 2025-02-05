package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secretkey      = []byte("secret-key")
	accesstokenexp = time.Minute * 15
	refreshtokexp  = time.Hour * 24 * 7
)

type Claims struct {
	Username string `json:"Username"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(username string) (string, error) {
	accessclaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accesstokenexp)),
			Issuer:    "https://example.com",
			Subject:   "access-token",
		},
	}
	accesstoken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessclaims)
	accesstokenstring, err := accesstoken.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return accesstokenstring, nil

}
func GenerateRefreshToken(username string) (string, error) {
	refreshclaims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshtokexp)),
			Issuer:    "https://example.com",
			Subject:   "refresh-token",
		},
	}
	refreshtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims)
	refreshtokenstring, err := refreshtoken.SignedString(secretkey)
	if err != nil {
		return "", err
	}
	return refreshtokenstring, nil
}

func validatetokens(tokenstring string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid  refresh token")
}
func GetExpiredRefershCookie() *http.Cookie {
	return &http.Cookie{

		Value:   "",
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
		Secure:  false,

		HttpOnly: true,
	}
}
