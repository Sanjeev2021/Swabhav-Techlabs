package contact

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"login/services"
)

// Contact represents the contact model
type Contact struct {
	ID           uint
	ContactName  string
	ContactType  string
	ContactValue string
	UserID       uint
	User         services.User
}

var database *gorm.DB

// InitDB initializes the database connection and performs necessary migrations
func InitDB() {
	// Database connection string
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

	err = database.AutoMigrate(&Contact{})
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

// CreateContact creates a new contact
func CreateContact(userid uint, contactname, contacttype string, contactvalue string) (*Contact, error) {
	db := GetDB()

	contact := &Contact{
		ContactName:  contactname,
		ContactValue: contactvalue,
		ContactType:  contacttype,
		UserID:       userid,
	}

	result := db.Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

// GetContactById retrieves a contact by its ID
func GetContactById(id uint) (*Contact, error) {
	db := GetDB()

	var contact Contact

	result := db.First(&contact, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contact, nil
}

// DeleteContact deletes a contact by its ID
func DeleteContact(id uint) (*Contact, error) {
	db := GetDB()

	contact, err := GetContactById(id)
	if err != nil {
		return nil, err
	}

	result := db.Unscoped().Delete(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}

// UpdateContact updates a contact's details
func UpdateContact(updatecontact *Contact) (*Contact, error) {
	db := GetDB()

	contact, err := GetContactById(updatecontact.ID)
	if err != nil {
		return nil, err
	}

	if updatecontact.ContactName != "" {
		contact.ContactName = updatecontact.ContactName
	}

	if updatecontact.ContactType != "" {
		contact.ContactType = updatecontact.ContactType
	}

	if updatecontact.ContactValue != "" {
		contact.ContactValue = updatecontact.ContactValue
	}

	result := db.Save(contact)
	if result.Error != nil {
		return nil, result.Error
	}

	return contact, nil
}

// FindAllContact retrieves all contacts with pagination
func FindAllContact(page, pagesize int) ([]*Contact, error) {
	db := GetDB()
	var contact []*Contact

	offset := (page - 1) * pagesize

	result := db.Offset(offset).Limit(pagesize).Find(&contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}
