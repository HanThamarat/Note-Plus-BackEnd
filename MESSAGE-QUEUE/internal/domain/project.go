package domain

import (
	"time"
	"gorm.io/gorm"
)

type Project struct {
	ID				uint			`grom:"primaryKey" json:"id"`
	Name 			string			`gorm:"type:varchar(100)" json:"name"`
	Description		string			`gorm:"type:varchar(100)" json:"description"`
	Status 			bool			`gorm:"default:true" json:"status"`
	CreatedAt		time.Time		`gorm:"column:created_at;" json:"created_at"`
	CreatedBy		uint			`gorm:"column:created_by;" json:"created_by"`
	OrgId			uint			`gorm:"column:org_id;" json:"org_id"`
  	UpdatedAt 		time.Time		`gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt	`gorm:"column:deleted_at;" json:"deleted_at"`
}

type ProjectDTO struct {
	Name			string				`json:"name"`
	Description		string				`json:"description"`
	Status 			bool				`json:"status"`	
	UserId			*uint				`json:"user_id"`
	OrgId			uint				`json:"org_id"`
}