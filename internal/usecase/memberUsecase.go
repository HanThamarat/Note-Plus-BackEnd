package usecase

import "github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"

type memberUsecase struct {
	repo domain.MemberRepository
}

func NewMemberUsecase(r domain.MemberRepository) domain.MemberUsecase {
	return &memberUsecase{
		repo: r,
	}
}

func (u *memberUsecase) CreateMember(dto domain.MemberDTO) (*domain.Member, error) {
	return u.repo.CreateMember(dto);
}

func (u *memberUsecase) FindInOrgMember(orgId int) (*domain.MemberResponse, error) {
	return u.repo.FindInOrgMember(orgId);
}