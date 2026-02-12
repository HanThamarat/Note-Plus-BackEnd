package domain

import (
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID 				uint				`grom:"primaryKey" json:"id"`
	OrgId			uint				`gorm:"column:orgId;" json:"orgId"`
	UserId			uint				`gorm:"column:userId;"json:"userId"`
	RoleId			uint				`gorm:"column:roleId;"json:"roleId"`
	CreatedAt		time.Time			`gorm:"column:created_at;" json:"created_at"`
  	UpdatedAt 		time.Time			`gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at;" json:"deleted_at"`
} 