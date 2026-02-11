package repository

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"gorm.io/gorm"
)

type gormOrgRepository struct {
	db *gorm.DB
}

func NewGormOrgRepository(db *gorm.DB) domain.OrgRepository {
	return &gormOrgRepository{db}
}

func (r *gormOrgRepository) CreateNewOrg(dto domain.OrgDTO) (*domain.Organizations, error) {
	var org domain.Organizations;

	org.Name 		= dto.Name;
	org.Description = dto.Description;
	org.Status 		= dto.Status;
	org.CreatedBy 	= *dto.UserId;

	result := r.db.Create(&org).Joins("organizations");

	if result.Error != nil {
		return  nil, result.Error;
	}

	return &org, nil;
}

func (r *gormOrgRepository) FindAllOrg() (*[]domain.Organizations, error) {
	var orgs []domain.Organizations;

	result := r.db.Find(&orgs).Where("status = ?", true);

	if result.Error != nil {
		return nil, result.Error;
	}

	return &orgs, nil;
}

func (r *gormOrgRepository) FindOrgById(orgId uint) (*domain.Organizations, error) {
	var org domain.Organizations;	
	result := r.db.First(&org, orgId);

	if result.Error != nil {
		return nil, result.Error;
	}

	return &org, nil;
}

func (r *gormOrgRepository) UpdateOrg(orgId uint, dto domain.OrgDTO) (*domain.Organizations, error) {

	org := domain.Organizations{ID: orgId}
	result := r.db.Model(&org).Updates(domain.Organizations{
		Name: dto.Name,
		Description: dto.Description,
		Status: dto.Status,
	});

	if result.Error != nil {
		return nil, result.Error;
	}

	return &org, nil;
}

func (r *gormOrgRepository) DeleteOrg(orgId uint) (*domain.Organizations, error) {
	var org domain.Organizations;	
	recheckId := r.db.First(&org, orgId);

	if recheckId.Error != nil {
		return nil, recheckId.Error;
	}

	result := r.db.Delete(&org);

	if result.Error != nil {
		return nil, result.Error;
	}

	return &org, nil;
}