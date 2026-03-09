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

type MemberDTO struct {
	OrgId			uint				`json:"orgId"`
	Identity		string				`json:"identity"`
	RoleId			uint				`json:"roleId"`
}

type UserQuery struct {
	ID 				uint				`json:"id"`
	Email			string				`json:"email"`
	Name 			string				`json:"name"`
	Username		string				`json:"username"`
}

type Userinfo struct {
	MemberId 		uint				`json:"memberId"`
	UserId			uint				`json:"userId"`
	Email			string				`json:"email"`
	Name 			string				`json:"name"`
	Role			string				`json:"role"`
}

type MemberResponse struct {
	Organization 	string				`json:"organization"`
	Member 			[]Userinfo			`json:"members"`
}

type UpdateMemberDTO struct {
	MemberId		uint				`json:"memberId"`
	RoleId			uint				`json:"roleId"`
}

type MemberRepository interface {
	CreateMember(dto MemberDTO) (*Member, error)
	FindInOrgMember(orgId int) (*MemberResponse, error)
	// UpdateOrgMember(dto UpdateMemberDTO) (*Member, error)
}

type MemberUsecase interface {
	CreateMember(dto MemberDTO) (*Member, error)
	FindInOrgMember(orgId int) (*MemberResponse, error)
	// UpdateOrgMember(dto UpdateMemberDTO) (*Member, error)
}