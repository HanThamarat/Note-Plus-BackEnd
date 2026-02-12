package domain

import (
	"time"

	"gorm.io/gorm"
)

type Organizations struct {
	ID 				uint				`grom:"primaryKey" json:"id"`
	Name			string				`gorm:"type:varchar(100)" json:"name"`
	Description		*string				`gorm:"type:varchar(100)" json:"description"`
	Status 			bool				`gorm:"default:true" json:"status"`
	CreatedAt		time.Time			`gorm:"column:created_at;" json:"created_at"`
	CreatedBy		uint				`gorm:"column:created_by;" json:"created_by"`
  	UpdatedAt 		time.Time			`gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at;" json:"deleted_at"`

	Member			[]Member			`gorm:"foreignKey:orgId;references:ID" json:"member"`
}

type OrgDTO struct {
	Name			string				`json:"name"`
	Description		string				`json:"description"`
	Status 			bool				`json:"status"`	
	UserId			*uint
}

type OrgRepository interface {
	CreateNewOrg(dto OrgDTO) (*Organizations, error)
	FindAllOrg() (*[]Organizations, error)
	FindOrgById(orgId uint) (*Organizations, error)
	UpdateOrg(orgId uint, dto OrgDTO) (*Organizations, error)
	DeleteOrg(orgId uint) (*Organizations, error)
}

type OrgUsecase interface {
	NewOrg(dto OrgDTO) (*Organizations, error)
	AllOrg() (*[]Organizations, error)
	OrgById(orgId uint) (*Organizations, error)
	UpdateOrg(orgId uint, dto OrgDTO) (*Organizations, error)
	DeleteOrg(orgId uint) (*Organizations, error)
}