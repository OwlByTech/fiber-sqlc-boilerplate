package handler

import (
	dto "optitech/internal/dto"
	cdto "optitech/internal/dto/client"
	"optitech/internal/service"

	"github.com/gofiber/fiber/v2"
)

func GetClientHandler(c *fiber.Ctx) error {
	params := c.AllParams()
	req := &cdto.GetClientReq{}
	if err := dto.ValidateParamsToDTO(params, req); err != nil {
		// TODO: create a error handling structure for bad and good message with
		// status, for example nestjs error handling in json form
		// you should implement the previous as a midddleware for errors
		// follow the next reference https://docs.gofiber.io/guide/error-handling/
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := service.GetClientService(*req)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(res)
}

func CreateClientHandler(c *fiber.Ctx) error {
	req := &cdto.CreateClientReq{}

	err := c.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = dto.ValidateDTO(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	r, err := service.CreateClientService(*req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.JSON(r)
}
