package service

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"bankingapp/errors"
	"bankingapp/models/account"
	"bankingapp/repository"
)

// AccountService Give Access to Update, Add, Delete Account
type AccountService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

// NewAccountService returns new instance of AccountService
func NewAccountService(db *gorm.DB, repo repository.Repository) *AccountService {
	return &AccountService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}

func (service *AccountService) CreateAccount(newAccount *account.Account) error {
	//  Creating unit of work.
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	// Add newAccount.
	err := service.repository.Add(uow, newAccount)
	if err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}

// GetAllAccounts returns all accounts
func (service *AccountService) GetAllAccounts(allAccounts *[]account.Account, totalCount *int) error {
	// Start new transcation.
	uow := repository.NewUnitOfWork(service.db, true)
	defer uow.RollBack()
	err := service.repository.GetAll(uow, allAccounts)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

// DOES ACCOUNT EXIST
func (service *AccountService) doesAccountExist(ID uint) error {
	exists, err := repository.DoesRecordExistForAccount(service.db, ID, account.Account{},
		repository.Filter("`id` = ?", ID))
	if !exists || err != nil {
		fmt.Println(">>>>>>>>>>>>>", exists)
		return errors.NewValidationError("Account does not exist ")
	}
	return nil
}

// UpdateAccount updates account
func (service *AccountService) UpdateAccount(accountToUpdate *account.Account) error {
	err := service.doesAccountExist(accountToUpdate.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	tempAccount := account.Account{}
	err = service.repository.GetRecordForAccount(uow, accountToUpdate.ID, &tempAccount, repository.Select("`created_at`"),
		repository.Filter("`id` = ?", accountToUpdate.ID))
	if err != nil {
		return err
	}
	accountToUpdate.CreatedAt = tempAccount.CreatedAt

	err = service.repository.Save(uow, accountToUpdate)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

// DeleteAccount deletes account
func (service *AccountService) DeleteAccount(accountToDelete *account.Account) error {
	err := service.doesAccountExist(accountToDelete.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()

	// Update test for updating deleted_by and deleted_at fields of test
	if err := service.repository.UpdateWithMap(uow, accountToDelete, map[string]interface{}{

		"DeletedAt": time.Now(),
	},
		repository.Filter("`id` = ?", accountToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}

	uow.Commit()
	return nil
}
