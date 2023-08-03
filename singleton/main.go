package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	genrateUID := uuid.New()

	alluser = append(alluser, User{
		Username: "yashshah",
		Password: "hello",
		UID:      genrateUID.String(),
		Role:     "admin",
		IsActive: "1",
		Fname:    "yash",
	})
	handleRequests()
}
func handleRequests() {
	headersOk := handlers.AllowCredentials()
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/login", loginPage)
	//router.HandleFunc("/refresh", refreshToken)
	router.HandleFunc("/api/v1/blog/users", getUser).Methods("GET")
	router.HandleFunc("/api/v1/blog/user", addUser).Methods("POST")
	router.HandleFunc("/api/v1/blog/deleteuser/{id}", deleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/blog/UpdateUser/{username}", UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/blog/getContact/{username}", getContact).Methods("GET")
	router.HandleFunc("/api/v1/blog/addContact/{username}", addContact).Methods("POST")
	router.HandleFunc("/api/v1/blog/getContactDetails/{username}/{Cname}", getContactDetails).Methods("GET")
	router.HandleFunc("/api/v1/blog/addContactDetails/{username}/{Cname}", addContactDetails).Methods("POST")
	log.Fatal(http.ListenAndServe(":4002", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	//refreshToken(w, r)
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
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
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
	json.NewEncoder(w).Encode(claims.Username)

}
func loginPage(w http.ResponseWriter, r *http.Request) {
	var userforlogin User
	var password string
	err := json.NewDecoder(r.Body).Decode(&userforlogin)
	password = userforlogin.Password
	fmt.Println("Inside Login")
	fmt.Println(userforlogin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if userforlogin.IsActive != "1" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	for _, u := range alluser {
		if u.Username == userforlogin.Username {
			userforlogin = u
			break
		}
	}
	fmt.Println(userforlogin)
	if password != userforlogin.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 10)
	claims := &Claims{
		Username: userforlogin.Username,
		Role:     userforlogin.Role,
		Fname:    userforlogin.Fname,
		UID:      userforlogin.UID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	fmt.Println(userforlogin.Role)
	json.NewEncoder(w).Encode(userforlogin.Role)

}
