package service

import (
	"time"

	"github.com/jinzhu/gorm"

	"bankingapp/errors"
	"bankingapp/models/user"
	"bankingapp/repository"
)

// UserService Give Access to Update, Add, Delete User
type UserService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

// NewUserService returns new instance of UserService
func NewUserService(db *gorm.DB, repo repository.Repository) *UserService {
	return &UserService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}
func (service *UserService) CreateUser(newUser *user.User) error {
	//  Creating unit of work.
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	// Add newUser.
	err := service.repository.Add(uow, newUser)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}
func (service *UserService) GetAllUsers(allUsers *[]user.User, totalCount *int) error {
	// Start new transcation.
	uow := repository.NewUnitOfWork(service.db, true)
	defer uow.RollBack()
	err := service.repository.GetAll(uow, allUsers)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
func (service *UserService) UpdateUser(userToUpdate *user.User) error {
	err := service.doesUserExist(userToUpdate.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	tempUser := user.User{}
	err = service.repository.GetRecordForUser(uow, userToUpdate.ID, &tempUser, repository.Select("`created_at`"),
		repository.Filter("`id` = ?", userToUpdate.ID))
	if err != nil {
		return err
	}
	userToUpdate.CreatedAt = tempUser.CreatedAt

	err = service.repository.Save(uow, userToUpdate)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}
func (service *UserService) doesUserExist(ID uint) error {
	exists, err := repository.DoesRecordExistForUser(service.db, ID, user.User{},
		repository.Filter("`id` = ?", ID))
	if !exists || err != nil {
		return errors.NewValidationError("User ID is Invalid")
	}
	return nil
}

func (service *UserService) DeleteUser(userToDelete *user.User) error {
	err := service.doesUserExist(userToDelete.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()

	// Update test for updating deleted_by and deleted_at fields of test
	if err := service.repository.UpdateWithMap(uow, userToDelete, map[string]interface{}{

		"DeletedAt": time.Now(),
	},
		repository.Filter("`id`=?", userToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}
