package service

import (
	"contactApp/repository"

	"github.com/jinzhu/gorm"
)

// ContactInfoService Give Access to Update, Add, Delete User
type ContactInfoService struct {
	db           *gorm.DB
	repository   repository.Repository
	associations []string
}

// NewContactInfoService returns new instance of ContactInfoService
func NewContactInfoService(db *gorm.DB, repo repository.Repository) *ContactInfoService {
	return &ContactInfoService{
		db:           db,
		repository:   repo,
		associations: []string{},
	}
}
