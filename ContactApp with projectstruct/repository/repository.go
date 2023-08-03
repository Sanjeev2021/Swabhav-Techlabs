package repository

import (
	"github.com/jinzhu/gorm"

	"contactApp/errors"
)

type Repository interface {
	GetAll(uow *UnitOfWork, out interface{}, queryProcessor ...QueryProcessor) error
	Add(uow *UnitOfWork, out interface{}) error
	Update(uow *UnitOfWork, out interface{}) error
	UpdateWithMap(uow *UnitOfWork, model interface{}, value map[string]interface{}, queryProcessors ...QueryProcessor) error
	GetRecordForUser(uow *UnitOfWork, userID uint, out interface{}, queryProcessors ...QueryProcessor) error
	Save(uow *UnitOfWork, value interface{}) error
}
type GormRepository struct{}

func NewGormRepository() *GormRepository {
	return &GormRepository{}
}

// UnitOfWork represent connection
type UnitOfWork struct {
	DB        *gorm.DB
	Committed bool
	Readonly  bool
}

// NewUnitOfWork creates new instance of UnitOfWork.
func NewUnitOfWork(db *gorm.DB, readonly bool) *UnitOfWork {
	commit := false
	if readonly {
		return &UnitOfWork{
			DB:        db.New(),
			Committed: commit,
			Readonly:  readonly,
		}
	}

	return &UnitOfWork{
		DB:        db.New().Begin(),
		Committed: commit,
		Readonly:  readonly,
	}
}

// Commit use to commit after a successful transaction.
func (uow *UnitOfWork) Commit() {
	if !uow.Readonly && !uow.Committed {
		uow.Committed = true
		uow.DB.Commit()
	}
}
// Save updates the record in table. If value doesn't have primary key, new record will be inserted.
func (repository *GormRepository) Save(uow *UnitOfWork, value interface{}) error {
	return uow.DB.Debug().Save(value).Error
}
// RollBack is used to rollback a transaction on failure.
func (uow *UnitOfWork) RollBack() {
	// This condition can be used if Rollback() is defered as soon as UOW is created.
	// So we only rollback if it's not committed.
	if !uow.Committed && !uow.Readonly {
		uow.DB.Rollback()
	}
}
func Filter(condition string, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Debug().Where(condition, args...)
		return db, nil
	}
}
// When creating/updating, specify fields that you want to save to database.
func Select(query interface{}, args ...interface{}) QueryProcessor {
	return func(db *gorm.DB, out interface{}) (*gorm.DB, error) {
		db = db.Select(query, args...)
		return db, nil
	}
}
// GetAll returns all records from the table.
func (repository *GormRepository) GetAll(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().Find(out).Error
}

// executeQueryProcessors executes all queryProcessor func.
func executeQueryProcessors(db *gorm.DB, out interface{}, queryProcessors ...QueryProcessor) (*gorm.DB, error) {
	var err error
	for _, query := range queryProcessors {
		if query != nil {
			db, err = query(db, out)
			if err != nil {
				return db, err
			}
		}
	}
	return db, nil
}

// Add adds record to table.
func (repository *GormRepository) Add(uow *UnitOfWork, out interface{}) error {
	return uow.DB.Create(out).Error
}

// Update updates the record in table.
func (repository *GormRepository) Update(uow *UnitOfWork, out interface{}) error {
	return uow.DB.Model(out).Update(out).Error
}
func DoesRecordExistForUser(db *gorm.DB, userID uint, out interface{}, queryProcessors ...QueryProcessor) (bool, error) {
	if userID == 0 {
		return false, errors.NewValidationError("DoesRecordExistForTenant: Invalid tenant ID")
	}
	count := 0
	// Below comment would make the tenant check before all query processor (Uncomment only if needed in future)
	// queryProcessors = append([]QueryProcessor{Filter("tenant_id = ?", tenantID)},queryProcessors... )
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return false, err
	}
	if err := db.Debug().Model(out).Where("id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
func (repository *GormRepository) GetRecordForUser(uow *UnitOfWork, userID uint, out interface{}, queryProcessors ...QueryProcessor) error {
	queryProcessors = append([]QueryProcessor{Filter("id = ?", userID)}, queryProcessors...)
	return repository.GetRecord(uow, out, queryProcessors...)
}
func (repository *GormRepository) GetRecord(uow *UnitOfWork, out interface{}, queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, out, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().First(out).Error
}
func (repository *GormRepository) UpdateWithMap(uow *UnitOfWork, model interface{}, value map[string]interface{},
	queryProcessors ...QueryProcessor) error {
	db := uow.DB
	db, err := executeQueryProcessors(db, value, queryProcessors...)
	if err != nil {
		return err
	}
	return db.Debug().Model(model).Update(value).Error
}