package user

import (
	"errors"
	//"fmt"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	IsAdmin   bool
	Username  string
	PassBook  []string
}

var UsersCreatedByMe []*User

func FindUser(userSlice []*User, id string) (*User, bool) {
	for i := 0; i < len(userSlice); i++ {
		if userSlice[i].ID.String() == id {
			return userSlice[i], true
		}
	}
	return nil, false
}

func DeleteUser(id string) (*User, error) {
	for i, user := range UsersCreatedByMe {
		if user.ID.String() == id {
			UsersCreatedByMe = append(UsersCreatedByMe[:i], UsersCreatedByMe[i+1:]...)
			return user, nil
		}
	}
	return nil, errors.New("no user found")
}

func UpdateUser(id, field, value string) (*User, error) {
	userToUpdate, userExist := FindUser(UsersCreatedByMe, id)
	if !userExist {
		return nil, errors.New("user does not exist")
	}
	switch field {
	case "FirstName":
		userToUpdate.FirstName = value
	case "LastName":
		userToUpdate.LastName = value
	case "Username":
		userToUpdate.Username = value
	}

	return userToUpdate, nil
}

func CreateUser(firstname, lastname, username string, isAdmin bool) (*User, error) {
	newUser := &User{
		ID:        uuid.NewV4(),
		FirstName: firstname,
		LastName:  lastname,
		IsAdmin:   isAdmin,
		Username:  username,
	}
	UsersCreatedByMe = append(UsersCreatedByMe, newUser)
	return newUser, nil
}
