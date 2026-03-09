package handler

import (
	"strconv"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type MemberHandler struct {
	usecase domain.MemberUsecase
}

func NewMemberHandler(uc domain.MemberUsecase) *MemberHandler {
	return &MemberHandler{
		usecase: uc,
	}
}

func (h *MemberHandler) CreateMember(c *fiber.Ctx) error {
	var req domain.MemberDTO;

	if err := c.BodyParser(&req); err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating member failed.",
			err.Error(),
		);
	}

	result, err := h.usecase.CreateMember(req);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating member failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusCreated,
		"member invited",
		result,
	);
}

func (h *MemberHandler) FindOrgMember(c *fiber.Ctx) error {
	orgId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Finding member by organization failed.",
			err.Error(),
		)
	}

	result, err := h.usecase.FindInOrgMember(orgId);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Finding member by organization failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Finding member by organization successfully.",
		result,
	);
}