package handler

import (
	dto "owlbytech/internal/dto"
	cdto "owlbytech/internal/dto/client"
	"owlbytech/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service interfaces.IClientService
}

func NewHandler(s interfaces.IClientService) interfaces.IClientHandler {
	return &handler{
		service: s,
	}
}

func (h *handler) Get(c *fiber.Ctx) error {
	data := c.Locals("clientId")
	clientId, ok := data.(int64)

	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot casting client id")
	}

	req := &cdto.GetClientReq{
		Id: clientId,
	}

	res, err := h.service.Get(req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func (h *handler) Create(c *fiber.Ctx) error {
	req := &cdto.CreateClientReq{}

	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	r, err := h.service.Create(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(r)
}

func (h *handler) Login(c *fiber.Ctx) error {
	req := &cdto.LoginClientReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.Login(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)

}

func (h *handler) ResetPassword(c *fiber.Ctx) error {
	req := &cdto.ResetPasswordReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.ResetPassword(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
func (h *handler) ResetPasswordToken(c *fiber.Ctx) error {
	req := &cdto.ResetPasswordTokenReq{}

	if err := c.BodyParser(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input: "+err.Error())
	}

	if err := dto.ValidateDTO(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := h.service.ResetPasswordToken(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
func (h *handler) ValidateResetPasswordToken(c *fiber.Ctx) error {
	token := c.Query("token")
	req := &cdto.ValidateResetPasswordTokenReq{
		Token: token,
	}
	res, err := h.service.ValidateResetPasswordToken(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}
