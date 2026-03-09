package usecase

import "github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"

type orgUsecase struct {
	repo domain.OrgRepository
}

func NewOrgUsecase(r domain.OrgRepository) domain.OrgUsecase {
	return &orgUsecase{repo: r}
}

func (u *orgUsecase) NewOrg(dto domain.OrgDTO) (*domain.Organizations, error) {
	return u.repo.CreateNewOrg(dto);
}

func (u *orgUsecase) AllOrg() (*[]domain.Organizations, error) {
	return u.repo.FindAllOrg();
}

func (u *orgUsecase) OrgById(orgId uint) (*domain.Organizations, error) {
	return u.repo.FindOrgById(orgId);
}

func (u *orgUsecase) UpdateOrg(orgId uint, dto domain.OrgDTO) (*domain.Organizations, error) {
	return u.repo.UpdateOrg(orgId, dto);
}

func (u *orgUsecase) DeleteOrg(orgId uint) (*domain.Organizations, error) {
	return  u.repo.DeleteOrg(orgId);
}

