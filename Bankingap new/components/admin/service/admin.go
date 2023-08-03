package service

import (
	"time"

	"github.com/jinzhu/gorm"

	"bankingapp/errors"
	"bankingapp/models/admin"
	"bankingapp/repository"
)

// UserService Give Access to Update, Add, Delete User
type AdminService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

// NewUserService returns new instance of UserService
func NewAdminService(db *gorm.DB, repo repository.Repository) *AdminService {
	return &AdminService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}
func (service *AdminService) CreateAdmin(newAdmin *admin.Admin) error {
	//  Creating unit of work.
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	// Add newUser.
	err := service.repository.Add(uow, newAdmin)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}
func (service *AdminService) GetAllAdmins(allAdmins *[]admin.Admin, totalCount *int) error {
	// Start new transcation.
	uow := repository.NewUnitOfWork(service.db, true)
	defer uow.RollBack()
	err := service.repository.GetAll(uow, allAdmins)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}
func (service *AdminService) UpdateAdmin(adminToUpdate *admin.Admin) error {
	err := service.doesAdminExist(adminToUpdate.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	tempAdmin := admin.Admin{}
	err = service.repository.GetRecordForAdmin(uow, adminToUpdate.ID, &tempAdmin, repository.Select("`created_at`"),
		repository.Filter("`id` = ?", adminToUpdate.ID))
	if err != nil {
		return err
	}
	adminToUpdate.CreatedAt = tempAdmin.CreatedAt

	err = service.repository.Save(uow, adminToUpdate)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}
func (service *AdminService) doesAdminExist(ID uint) error {
	exists, err := repository.DoesRecordExistForAdmin(service.db, ID, admin.Admin{},
		repository.Filter("`id` = ?", ID))
	if !exists || err != nil {
		return errors.NewValidationError("User ID is Invalid")
	}
	return nil
}

func (service *AdminService) DeleteAdmin(adminToDelete *admin.Admin) error {
	err := service.doesAdminExist(adminToDelete.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()

	// Update test for updating deleted_by and deleted_at fields of test
	if err := service.repository.UpdateWithMap(uow, adminToDelete, map[string]interface{}{

		"DeletedAt": time.Now(),
	},
		repository.Filter("`id`=?", adminToDelete)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}
