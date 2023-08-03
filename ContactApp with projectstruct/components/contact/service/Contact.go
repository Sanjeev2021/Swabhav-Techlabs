package service

import (
	"contactApp/repository"

	"github.com/jinzhu/gorm"
)

// ContactService Give Access to Update, Add, Delete Contact
type ContactService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

// NewContactService returns new instance of ContactService
func NewContactService(db *gorm.DB, repo repository.Repository) *ContactService {
	return &ContactService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}
