package handler

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	usecase domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: uc,
	}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var req domain.UserDTO;
	if err := c.BodyParser(&req); err != nil {
		return responses.SetErrResponse(c, fiber.StatusBadRequest, "Creating a user failed.", err.Error());
	}

	result, err := h.usecase.Register(req);

	if err != nil {
		return responses.SetErrResponse(c, fiber.StatusBadRequest, "Creating a user failed.", err.Error());
	}

	return responses.SetResponse(
		c,
		fiber.StatusCreated,
		"Creating a user successfully.",
		result,
	);
}