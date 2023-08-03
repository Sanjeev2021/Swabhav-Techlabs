package service

import (
	"time"

	"github.com/jinzhu/gorm"

	"bankingapp/errors"
	"bankingapp/models/admin"
	"bankingapp/models/bank"
	"bankingapp/repository"
)

type BankService struct {
	db         *gorm.DB
	repository repository.Repository
}

func NewBankService(db *gorm.DB, repository repository.Repository) *BankService {
	return &BankService{
		db:         db,
		repository: repository,
	}
}

func (service *BankService) CreateBank(newBank *bank.Bank) error {
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()
	err := service.repository.Add(uow, newBank)
	if err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}

func (service *BankService) GetAllBanks(allBanks *[]bank.Bank, totalCount *int) error {
	uow := repository.NewUnitOfWork(service.db, true)
	defer uow.RollBack()
	err := service.repository.GetAll(uow, allBanks)
	if err != nil {
		return err
	}
	uow.Commit()
	return nil
}

func (service *BankService) doesBankExist(Id uint) error {
	exists, err := repository.DoesRecordExistForBank(service.db, Id, admin.Admin{},
		repository.Filter("`id` = ?", Id))
	if err != nil {
		return err
	}
	if !exists {
		return errors.NewValidationError("Bank does not exist")
	}
	return nil

}
func (service *BankService) UpdateBank(bankToUpdate *bank.Bank) error {
	err := service.doesBankExist(bankToUpdate.ID)
	if err != nil {
		return err
	}

	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()

	// Define query processors for the GetRecordForBank method.
	selectProcessor := repository.Select("`created_at`")

	tempBank := bank.Bank{}
	err = service.repository.GetRecordForBank(uow, bankToUpdate.ID, &tempBank, selectProcessor)
	if err != nil {
		return err
	}

	bankToUpdate.CreatedAt = tempBank.CreatedAt

	err = service.repository.Save(uow, bankToUpdate)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}

func (service *BankService) DeleteBank(bankToDelete *bank.Bank) error {
	err := service.doesBankExist(bankToDelete.ID)
	if err != nil {
		return err
	}
	uow := repository.NewUnitOfWork(service.db, false)
	defer uow.RollBack()

	if err := service.repository.UpdateWithMap(uow, bankToDelete, map[string]interface{}{
		"DeletedAt": time.Now(),
	},
		repository.Filter("`id` = ?", bankToDelete.ID)); err != nil {
		uow.RollBack()
		return err
	}
	uow.Commit()
	return nil
}
