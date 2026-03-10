package usecase

import "github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"

type projectUsecase struct {
	repo domain.ProjectRepository
}

func NewProjectUsecase(r domain.ProjectRepository) domain.ProjectUsecase {
	return  &projectUsecase{repo: r}
}

func (u *projectUsecase) CreateProject(dto domain.ProjectDTO) (*domain.Project, error) {
	return u.repo.CreateProject(dto);
}

func (u *projectUsecase) FindAllOrgProject(orgId int) (*[]domain.Project, error) {
	return u.repo.FindAllOrgProject(orgId);
}