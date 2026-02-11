package repository

import (
	"errors"
	"os"
	"time"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/encrypt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type gormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &gormAuthRepository{db}
}

func (r *gormAuthRepository) CreateCredentialAuth(authDTO domain.AuthDTO) (*domain.AuthEntity, error) {
	var user domain.User;

	errs := r.db.Where("(email = ? OR username = ?) AND status = ?", authDTO.Username, authDTO.Username, true).First(&user).Error;
	if errs != nil {
		return nil, errors.New("This username or password not have in the system.");
	}

	verifyPassword := encrypt.VerifyPassword(authDTO.Password, *user.Password);

	if !verifyPassword {
		return nil, errors.New("Please check your password.");
	}

	claims := jwt.MapClaims{
		"userId": user.ID,
		"name": user.Name,
		"email": user.Email,
		"status": user.Status,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, errors.New("Have something wrong in generate jwt token, Please try again later.");
	}

	authEntity := domain.AuthEntity{
		Email: user.Email,
		Name: user.Name,
		AuthToken: t,
	}

	return &authEntity, nil;
}