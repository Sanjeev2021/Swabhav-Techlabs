package services

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"bankingApp/database"
)

// User represents the user model
type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Password  string
	Username  string `gorm:"unique"`
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &User{})
}

// CreateUser creates a new user
func CreateUser(firstname, lastname, username, password string) (*User, error) {
	db := database.GetDB()

	password, err := HashPassword(password) // here we are calling hash password
	if err != nil {
		return nil, errors.New("unable to hash password")
	}

	user := &User{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Password:  password,
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserById retrieves a user by its ID
func GetUserById(id uint) (*User, error) {
	db := database.GetDB()
	var user User

	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser deletes a user by its ID
func DeleteUser(id uint) (*User, error) {
	db := database.GetDB()

	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	result := db.Delete(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates a user's details
func UpdateUser(id uint, updateduser *User) (*User, error) {
	db := database.GetDB()

	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	if updateduser.FirstName != "" {
		user.FirstName = updateduser.FirstName
	}

	if updateduser.LastName != "" {
		user.LastName = updateduser.LastName
	}

	if updateduser.Password != "" {
		password, err := HashPassword(updateduser.Password) // here as well
		if err != nil {
			return nil, errors.New("unable to hash password")
		}
		user.Password = password
	}

	if updateduser.Username != "" {
		user.Username = updateduser.Username
	}

	result := db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// GetUserByUsername retrieves a user by its username
func GetUserByUsername(username string) (*User, error) {
	db := database.GetDB()
	var user User

	result := db.First(&user, "username = ?", username) // we are getting the user by username
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(page, pagesize int) ([]*User, error) {
	db := database.GetDB()
	var users []*User
	offset := (page - 1) * pagesize // page = current page no , pahesize = no of records to display per page , offset = no of records to skip

	result := db.Offset(offset).Limit(pagesize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // array of byte
	if err != nil {
		return "", err
	}
	return string(pass), nil
}
