package handler

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	usecase domain.AuthUsecase
}

func NewAuthHandler(uc domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		usecase: uc,
	}
}

func (h *AuthHandler) CredentialAuthenticate(c *fiber.Ctx) error {
	var reqs domain.AuthDTO;
	if err := c.BodyParser(&reqs); err != nil {
		return responses.SetErrResponse(
			c, 
			fiber.StatusBadRequest, 
			"Authentication failed.", 
			err.Error(),
		);
	}

	result, err := h.usecase.CredentialAuth(reqs);

	if err != nil {
		return responses.SetErrResponse(
			c, 
			fiber.StatusBadRequest, 
			"Authentication failed.", 
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Authentication successfully.",
		result,
	);
}