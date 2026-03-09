package domain

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID 				uint				`grom:"primaryKey" json:"id"`
	Name			string				`gorm:"type:varchar(100)" json:"name"`
	Description		*string				`gorm:"type:varchar(100)" json:"description"`
	Status 			bool				`gorm:"default:true" json:"status"`
	CreatedAt		time.Time			`gorm:"column:created_at;" json:"created_at"`
  	UpdatedAt 		time.Time			`gorm:"column:updated_at;" json:"updated_at"`
	DeletedAt 		gorm.DeletedAt		`gorm:"column:deleted_at;" json:"deleted_at"`

	Member			[]Member			`gorm:"foreignKey:roleId;references:ID" json:"member"`
}