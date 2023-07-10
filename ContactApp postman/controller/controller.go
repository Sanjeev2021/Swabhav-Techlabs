package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	user "contactapp/Services"
	"contactapp/contact"
	contactinfo "contactapp/contactInfo"
)

// add defer in every controller

type UpdateUser struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type ContactUpdate struct {
	ContactID string
	Field     string
	Value     string
}

type ContactDelete struct {
	ContactID string
}

type ContactInfoUpdate struct {
	ContactID string
	Field     string
	Value     string
}

type DeletedContactInfo struct {
	ContactID string
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing information", http.StatusBadRequest)
		return
	}

	userobj, err := user.FindUserById(id)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(user)
	var updatedUser *UpdateUser
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		fmt.Println(err)
	}

	// Perform the update operation on the user object
	updatedUserObj, err := userobj.UpdatedUser(id, updatedUser.Field, updatedUser.Value)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(updatedUserObj)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		http.Error(w, "Missing information", http.StatusBadRequest)
		return
	}

	user, err := user.FindUserById(id)
	if err != nil {
		fmt.Println(err)
	}

	err = user.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(&user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *user.User
	// r.Body means request body

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	newUser, err = user.CreateUser(newUser.FirstName, newUser.LastName, newUser.IsAdmin, newUser.Username)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(&newUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Missing information", http.StatusBadRequest)
		return
	}
	user, err := user.FindUserById(id)
	if err != nil {
		fmt.Println("err: ", err)
	}
	json.NewEncoder(w).Encode(&user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users *[]user.User
	users, err := user.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(&users)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userfound, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
	var newContact *contact.Contact
	// difference between Decoder and NewDecoder
	err = json.NewDecoder(r.Body).Decode(&newContact)
	if err != nil {
		http.Error(w, "missing parameters hai boss", http.StatusNotFound)
		return
	}

	newContact = userfound.CreateContact(newContact.ContactName, newContact.ContactType, newContact.ContactValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newContact)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	contact, err := user.FindContactById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&contact)

}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	founduser, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updateContact *ContactUpdate
	err = json.NewDecoder(r.Body).Decode(&updateContact)
	if err != nil {
		http.Error(w, "Missing paramater", http.StatusBadRequest)
		return
	}

	updatedContact, err := founduser.UpdateContact(updateContact.ContactID, updateContact.Field, updateContact.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&updatedContact)

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	finduser, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var decodedcontact *ContactDelete
	err = json.NewDecoder(r.Body).Decode(&decodedcontact)
	if err != nil {
		http.Error(w, "Missing parameter", http.StatusBadRequest)
		return
	}

	deltedcontact, err := finduser.DeleteContact(decodedcontact.ContactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&deltedcontact)
}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	var contacts *[]contact.Contact
	contacts, err := user.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&contacts)
}

func CreateContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userfound, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var newContactInfo *contactinfo.ContactInfo
	err = json.NewDecoder(r.Body).Decode(&newContactInfo)
	if err != nil {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}

	newContactInfo, err = userfound.CreateContactInfo(newContactInfo.ContactInfoType, newContactInfo.ContactInfoValue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&newContactInfo)
}

func UpdateContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	founduser, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var UpdateContactInfo *ContactInfoUpdate
	err = json.NewDecoder(r.Body).Decode(&UpdateContactInfo)
	if err != nil {
		http.Error(w, "Missing paramter", http.StatusBadRequest)
	}

	UpdatedContactInfo, err := founduser.UpdateContactInfo(UpdateContactInfo.ContactID, UpdateContactInfo.Field, UpdateContactInfo.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&UpdatedContactInfo)
}

func DeleteContactInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	deleteinfocontact, err := user.FindUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var deleteinfo *DeletedContactInfo
	err = json.NewDecoder(r.Body).Decode(&deleteinfo)
	if err != nil {
		http.Error(w, "missing paramter", http.StatusBadRequest)
	}

	deletedinfo, err := deleteinfocontact.DeleteContactInfo(deleteinfo.ContactID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&deletedinfo)
}

// func GetContactById(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	contact, err := user.FindContactById(id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(&contact)

// }

func GetAllContactInfo(w http.ResponseWriter, r *http.Request) {
	var contactinfo *[]contactinfo.ContactInfo
	contactinfo, err := user.GetAllContactInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&contactinfo)
}

func GetContactInfoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	contactinfo, err := user.FindContactInfoById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&contactinfo)

}
