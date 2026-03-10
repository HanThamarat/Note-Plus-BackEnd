package handler

import (
	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/encrypt"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type ProjectHandler struct {
	usecase domain.ProjectUsecase
}

func NewProjectHandler(uc domain.ProjectUsecase) *ProjectHandler {
	return &ProjectHandler{
		usecase: uc,
	}
}

func (h *ProjectHandler) CreateNewProject(c *fiber.Ctx) error {
	var req domain.ProjectDTO;
	if err := c.BodyParser(&req); err != nil {
		return responses.SetErrResponse(
			c,
			422,
			"Creating a new project failed.",
			err.Error(),
		);
	}

	userInfo, err := encrypt.JWTDecryption(c);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating a new project failed.",
			err.Error(),
		);
	}

	idUint := uint(userInfo.UserId);
	req.UserId = &idUint;

	result, err := h.usecase.CreateProject(req);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating a new project failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusCreated,
		"Creating a new project successfully.",
		result,
	);
}