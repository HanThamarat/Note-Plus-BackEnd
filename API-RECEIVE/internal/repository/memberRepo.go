package repository

import (
	"errors"
	"strings"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"gorm.io/gorm"
)

type gormMemberRepository struct {
	db	*gorm.DB
}

func NewGormMemberRepository(db	*gorm.DB) domain.MemberRepository {
	return &gormMemberRepository{db}
}

func (r *gormMemberRepository) CreateMember(dto domain.MemberDTO) (*domain.Member, error) {
	var userModel domain.User;
	var userQuery domain.UserQuery;

	err := r.db.Model(&userModel).
	Where("LOWER(email) = ? OR LOWER(username) = ?", strings.ToLower(dto.Identity), strings.ToLower(dto.Identity)).
	First(&userQuery).Error; 

	if err != nil {
		return nil, err;
	}

	var member domain.Member;
	r.db.Where(`"orgId" = ? AND "userId" = ?`,dto.OrgId, userQuery.ID).First(&member);

	if member.ID != 0 {
		return nil, errors.New("This user already have in organization");
	}

	newObject := domain.Member{
		OrgId: dto.OrgId,
		UserId: userQuery.ID,
		RoleId: dto.RoleId,
	};

	if err := r.db.Create(&newObject).Error; err != nil {
		return nil, err;
	}

	return &newObject, nil;
}

func (r *gormMemberRepository) FindInOrgMember(orgId int) (*domain.MemberResponse, error) {
	var org domain.Organizations;

	if err := r.db.Select("name").Where(`"id" = ?`, orgId).First(&org).Error; err != nil {
		return nil, err;
	}

	var member 		domain.Member;
	var userInfo 	[]domain.Userinfo;

	if err := r.db.Model(&member).
	Select(`u."name", u."email", r."name" as role, u."id" as "userId", members.id as "memberId"`).
	Where(`"orgId" = ?`, orgId).
	Joins(`INNER JOIN users as u on members."userId" = u.id`).
	Joins(`INNER JOIN roles as r on members."roleId" = r.id`).
	Scan(&userInfo).Error; err != nil {
		return nil, err;
	}
	
	result := domain.MemberResponse{
		Organization: org.Name,
		Member: userInfo,
	}	

	return &result, nil;
}

// func (r *gormMemberRepository) UpdateOrgMember(dto domain.UpdateMemberDTO) (*domain.Member, error) {

// }