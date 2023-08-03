package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside getContact")
	if isValidCoockie(w, r) {
		var contactsfordisplay []Contact
		var user User
		params := mux.Vars(r)
		for _, u := range alluser {
			if u.Username == params[("username")] {
				user = u
				break
			}
		}
		if user.Role != "admin" {
			for _, c := range allContact {
				if c.UID == user.UID {
					contactsfordisplay = append(contactsfordisplay, c)

				}
			}
			json.NewEncoder(w).Encode(contactsfordisplay)
			w.WriteHeader(200)
		} else {
			json.NewEncoder(w).Encode(allContact)
			w.WriteHeader(200)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}
}

func addContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside getContact")
	if isValidCoockie(w, r) {
		var newcontact Contact
		_ = json.NewDecoder(r.Body).Decode(&newcontact)
		params := mux.Vars(r)
		for _, u := range alluser {
			if u.Username == params[("username")] {
				newcontact.UID = u.UID
				break
			}
		}
		if newcontact.UID == "" {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Username not found")
		}
		var genrateCID = uuid.New()
		newcontact.CID = genrateCID.String()
		fmt.Println(newcontact)
		allContact = append(allContact, newcontact)
		json.NewEncoder(w).Encode(newcontact)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}
}
