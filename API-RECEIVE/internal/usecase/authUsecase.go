package usecase

import "github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"

type authUsecase struct {
	repo domain.AuthRepository
}

func NewAuthUsecase(r domain.AuthRepository) domain.AuthUsecase {
	return &authUsecase{repo: r}
}

func (u *authUsecase) CredentialAuth(authDTO domain.AuthDTO) (*domain.AuthEntity, error) {
	return u.repo.CreateCredentialAuth(authDTO);
}