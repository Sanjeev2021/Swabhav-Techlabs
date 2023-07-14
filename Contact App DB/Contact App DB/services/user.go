package services

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User represents the user model
type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Password  string
	Username  string `gorm:"unique"`
	CreatedAt time.Time
}

var database *gorm.DB

// InitDB initializes the database connection and performs necessary migrations
func InitDB() {
	// Database connection string , dsn = data source name
	dsn := "root:root@tcp(127.0.0.1:3306)/contactapp?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	fmt.Println("connection established")

	sqlDB, err := database.DB()
	if err != nil {
		panic("failed to access underlying DB object: " + err.Error())
	}

	err = sqlDB.Ping()
	if err != nil {
		panic("failed to ping database: " + err.Error())
	}

	err = database.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database :" + err.Error())
	}
	fmt.Println("database migrated")

	fmt.Println("MySQL connection established!")
}

// GetDB returns the database connection instance
func GetDB() *gorm.DB {
	if database == nil {
		InitDB()
	}

	return database
}

// CreateUser creates a new user
func CreateUser(firstname, lastname, username, password string) (*User, error) {
	db := GetDB()
	user := &User{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// GetUserById retrieves a user by its ID
func GetUserById(id uint) (*User, error) {
	db := GetDB()
	var user User

	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser deletes a user by its ID
func DeleteUser(id uint) (*User, error) {
	db := GetDB()

	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}
	// here it is doing hard delete
	result := db.Unscoped().Delete(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates a user's details
func UpdateUser(updateduser *User) (*User, error) {
	db := GetDB()

	user, err := GetUserById(updateduser.ID)
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
		user.Password = updateduser.Password
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
	db := GetDB()
	var user User

	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(page, pagesize int) ([]*User, error) {
	db := GetDB()
	var users []*User
	offset := (page - 1) * pagesize // page = current page no , pahesize = no of records to display per page , offset = no of records to skip

	result := db.Offset(offset).Limit(pagesize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
