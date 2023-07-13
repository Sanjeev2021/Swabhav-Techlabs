package contactinfo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ContactInfo represents the contact information model
type ContactInfo struct {
	ID               uint
	UserID           uint
	ContactInfoType  string
	ContactInfoValue string
}

// CreateContactInfo creates a new contact information entry
func CreateContactInfo(userid uint, contactinfotype, contactinfovalue string) (*ContactInfo, error) {
	db := GetDB()

	contactinfo := &ContactInfo{
		UserID:           userid,
		ContactInfoType:  contactinfotype,
		ContactInfoValue: contactinfovalue,
	}

	result := db.Create(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
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

	err = database.AutoMigrate(&ContactInfo{})
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

// GetContactInfoById retrieves a contact information entry by its ID
func GetContactInfoById(id uint) (*ContactInfo, error) {
	db := GetDB()

	var contactinfo ContactInfo

	result := db.First(&contactinfo, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &contactinfo, nil
}

// DeleteContactInfo deletes a contact information entry by its ID
func DeleteContactInfo(id uint) (*ContactInfo, error) {
	db := GetDB()

	contactinfo, err := GetContactInfoById(id)
	if err != nil {
		return nil, err
	}

	result := db.Unscoped().Delete(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
}

// UpdateContactInfo updates a contact information entry's details
func UpdateContactInfo(updatecontactinfo *ContactInfo) (*ContactInfo, error) {
	db := GetDB()

	contactinfo, err := GetContactInfoById(updatecontactinfo.ID)
	if err != nil {
		return nil, err
	}

	if updatecontactinfo.ContactInfoType != "" {
		contactinfo.ContactInfoType = updatecontactinfo.ContactInfoType
	}

	if updatecontactinfo.ContactInfoValue != "" {
		contactinfo.ContactInfoValue = updatecontactinfo.ContactInfoValue
	}

	result := db.Save(contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}

	return contactinfo, nil
}

// FindAllContactInfo retrieves all contact information entries with pagination
func FindAllContactInfo(page, pagesize int) ([]*ContactInfo, error) {
	db := GetDB()

	var contactinfo []*ContactInfo

	offset := (page - 1) * pagesize

	result := db.Offset(offset).Limit(page).Find(&contactinfo)
	if result.Error != nil {
		return nil, result.Error
	}
	return contactinfo, nil
}
