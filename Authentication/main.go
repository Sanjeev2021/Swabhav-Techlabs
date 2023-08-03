package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/refresh", Refresh)
	// no router and framework so nil
	log.Fatal(http.ListenAndServe(":8080", nil))
}
