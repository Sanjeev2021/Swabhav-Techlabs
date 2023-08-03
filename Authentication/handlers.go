package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "github.com/golang-jwt/jwt/v4"
	"github.com/dgrijalva/jwt-go"
)

// creating jwt key
var jwtKey = []byte("secret_key")

var users = map[string]string{
	// key value pair
	// we can use db as well
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Username string `json:"username`
	Password string `json: "password"`
}

// claim struct provides expiration time and issue time
type Claims struct {
	Username string `json:"username`
	jwt.StandardClaims
}

// responsewriter used to write the response back to client
func Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials                         // declaring variable name of type credential
	err := json.NewDecoder(r.Body).Decode(&credentials) // r.Body is basically request body , decode will decode the data and store in credential var
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]  //matching username & password
	if !ok || expectedPassword != credentials.Password { // if data not available or password not match
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5) // adding 5 minutes

	claims := &Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //new jwt token is being created with signingmethodhs256
	tokenString, err := token.SignedString(jwtKey)             //signed using jwt key
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

// http request it contains information about the incoming http request, like
// url , headers
func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Username)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5) // adding 5 minutes

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //new jwt token is being created with signingmethodhs256
	tokenString, err := token.SignedString(jwtKey)             //signed using jwt key
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   tokenString,
		Expires: expirationTime,
	})

}

