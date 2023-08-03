package Admin

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"bankingApp/database"
)

type Admin struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Password  string
	Username  string `gorm:"unique"`
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &Admin{})
}

func CreateAdmin(firstname, lastname, username, password string) (*Admin, error) {
	db := database.GetDB()

	password, err := HashPassword(password) // here we are calling hashpassword
	if err != nil {
		return nil, errors.New("unable to hash password")
	}

	Admin := &Admin{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Password:  password,
	}

	result := db.Create(Admin)
	if result.Error != nil {
		return nil, result.Error
	}
	return Admin, nil
}

// GetAdmin by id
func GetAdminById(id uint) (*Admin, error) {
	db := database.GetDB()
	var admin Admin

	result := db.First(&admin, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &admin, nil
}

func DeleteAdmin(id uint) (*Admin, error) {
	db := database.GetDB()

	admin, err := GetAdminById(id)
	if err != nil {
		return nil, err
	}

	result := db.Delete(admin)
	if result.Error != nil {
		return nil, result.Error

	}
	return admin, nil
}

func UpdateAdmin(id uint, updateadmin *Admin) (*Admin, error) {
	db := database.GetDB()

	admin, err := GetAdminById(id)
	if err != nil {
		return nil, err
	}

	if updateadmin.FirstName != "" {
		admin.FirstName = updateadmin.FirstName
	}

	if updateadmin.LastName != "" {
		admin.LastName = updateadmin.LastName
	}

	if updateadmin.Password != "" {
		password, err := HashPassword(updateadmin.Password)
		if err != nil {
			return nil, errors.New("unable to hash password")
		}
		admin.Password = password
	}

	if updateadmin.Username != "" {
		admin.Username = updateadmin.Username
	}

	result := db.Save(admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}

// ye hashpassword call kaha pe kar rhe hai ?
func HashPassword(password string) (string, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func GetAllAdmins(page, pagesize int) ([]*Admin, error) {
	db := database.GetDB()
	var users []*Admin
	offset := (page - 1) * pagesize // page = current page no , pahesize = no of records to display per page , offset = no of records to skip

	result := db.Offset(offset).Limit(pagesize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
