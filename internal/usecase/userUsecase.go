package usecase

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/encrypt"
)

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: r}
}

func (u *userUsecase) Register(userDTO domain.UserDTO) (*domain.User, error) {

	hash, _ := encrypt.HashPassword(userDTO.Password);

	user := &domain.User{
		Email: userDTO.Email,
		Name: userDTO.Name,
		Username: userDTO.Username,
		Password: &hash,
		Status: userDTO.Status,
	}

	return u.repo.Create(user);
}