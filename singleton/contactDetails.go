package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func getContactDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside getContact")
	if isValidCoockie(w, r) {
		var contactdetailsfordisplay []ContactDetails
		var contact Contact
		params := mux.Vars(r)
		var uid = ""
		for _, u := range alluser {
			if u.Username == params[("username")] {
				uid = u.UID
				break
			}
		}
		if uid == "" {
			w.WriteHeader(409)
			fmt.Fprintf(w, "User not found")
		}
		var contactsofuser []Contact
		for _, c := range allContact {
			if c.UID == uid {
				contactsofuser = append(contactsofuser, c)
			}
		}
		for _, c := range contactsofuser {
			if c.Cname == params[("Cname")] {
				contact = c
				break
			}
		}

		for _, cd := range allContactdetails {
			if cd.CID == contact.CID {
				contactdetailsfordisplay = append(contactdetailsfordisplay, cd)

			}
		}
		json.NewEncoder(w).Encode(contactdetailsfordisplay)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}
}

func addContactDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Println("Inside getContact")
	if isValidCoockie(w, r) {
		var newcontactdetail ContactDetails
		_ = json.NewDecoder(r.Body).Decode(&newcontactdetail)
		params := mux.Vars(r)

		var uid = ""
		for _, u := range alluser {
			if u.Username == params[("username")] {
				uid = u.UID
				break
			}
		}
		if uid == "" {
			w.WriteHeader(409)
			fmt.Fprintf(w, "User not found")
		}
		var contactsofuser []Contact
		for _, c := range allContact {
			if c.UID == uid {
				contactsofuser = append(contactsofuser, c)
			}
		}
		for _, c := range contactsofuser {
			if c.Cname == params[("Cname")] {
				newcontactdetail.CID = c.CID
				break
			}
		}
		if newcontactdetail.CID == "" {
			w.WriteHeader(409)
			fmt.Fprintf(w, "Contact not found")
		}
		var genrateCDID = uuid.New()
		newcontactdetail.CDID = genrateCDID.String()
		fmt.Println(newcontactdetail)
		allContactdetails = append(allContactdetails, newcontactdetail)
		json.NewEncoder(w).Encode(newcontactdetail)
		w.WriteHeader(200)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Unauthorised Access")
	}
}
