package handler

import (
	"strconv"

	"github.com/HanThamarat/Note-Plus-BackEnd/internal/domain"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/encrypt"
	"github.com/HanThamarat/Note-Plus-BackEnd/pkg/responses"
	"github.com/gofiber/fiber/v2"
)

type OrgHandler struct {
	usecase domain.OrgUsecase
}

func NewOrgHandler(uc domain.OrgUsecase) *OrgHandler {
	return &OrgHandler{
		usecase: uc,
	}
} 

func (h *OrgHandler) CreateNewOrg(c *fiber.Ctx) error {
	var req domain.OrgDTO;
	if err := c.BodyParser(&req); err != nil {
		return responses.SetErrResponse(
			c,
			422,
			"Creating a new organization failed.",
			err.Error(),
		);
	}

	userInfo, err := encrypt.JWTDecryption(c);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating a new organization failed.",
			err.Error(),
		);
	}

	idUint := uint(userInfo.UserId);
	req.UserId = &idUint;

	result, err := h.usecase.NewOrg(req);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating a new organization failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusCreated,
		"Creating a new organization successfully.",
		result,
	);
}

func (h *OrgHandler) FindAllOrg(c *fiber.Ctx) error {
	orgs, err := h.usecase.AllOrg();

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Find all organization failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Find all organization successfully.",
		orgs,
	);
}

func (h *OrgHandler) FindOrgById(c *fiber.Ctx) error {
	orgId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Find organization by id failed.",
			err.Error(),
		);
	}

	result, err := h.usecase.OrgById(uint(orgId));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Find organization by id failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Find organization by id successfully.",
		result,
	);
}

func (h *OrgHandler) UpdateOrg(c *fiber.Ctx) error {
	var req domain.OrgDTO;
	orgId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Update organization by id failed.",
			err.Error(),
		);
	}

	if err := c.BodyParser(&req); err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Update organization by id failed.",
			err.Error(),
		);
	}

	userInfo, err := encrypt.JWTDecryption(c);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Creating a new organization failed.",
			err.Error(),
		);
	}

	idUint := uint(userInfo.UserId);
	req.UserId = &idUint;

	result, err := h.usecase.UpdateOrg(uint(orgId), req);

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Update organization by id failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Update organization by id successfully.",
		result,
	);
}

func (h *OrgHandler) DeleteOrg(c *fiber.Ctx) error {
	orgId, err := strconv.Atoi(c.Params("id"));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Delete organization by id failed.",
			err.Error(),
		);
	}

	result, err := h.usecase.DeleteOrg(uint(orgId));

	if err != nil {
		return responses.SetErrResponse(
			c,
			fiber.StatusBadRequest,
			"Delete organization by id failed.",
			err.Error(),
		);
	}

	return responses.SetResponse(
		c,
		fiber.StatusOK,
		"Delete organization by id successfully.",
		result,
	);
}