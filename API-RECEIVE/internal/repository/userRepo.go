package repository

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) domain.UserRepository {
	return &gormUserRepository{db}
}

func (r *gormUserRepository) Create(user *domain.User) (*domain.User, error) {
	
	result := r.db.Create(user);

	if result.Error != nil {
		return nil, result.Error;
	}

	return user, nil;
}