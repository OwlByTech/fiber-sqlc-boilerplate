package interfaces

import (
	dto "owlbytech/internal/dto/client"
	sq "owlbytech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IClientHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	ResetPassword(c *fiber.Ctx) error
	ResetPasswordToken(c *fiber.Ctx) error
	ValidateResetPasswordToken(c *fiber.Ctx) error
}

type IClientService interface {
	Get(req *dto.GetClientReq) (*dto.GetClientRes, error)
	Create(req *dto.CreateClientReq) (*dto.CreateClientRes, error)
	UpdateById(req *dto.UpdateClientByIdReq) (bool, error)
	Login(req *dto.LoginClientReq) (*dto.LoginClientRes, error)
	ResetPassword(req dto.ResetPasswordReq) (bool, error)
	ResetPasswordToken(req *dto.ResetPasswordTokenReq) (bool, error)
	ValidateResetPasswordToken(req dto.ValidateResetPasswordTokenReq) (bool, error)
}

type IClientRepository interface {
	Get(id *int64) (*sq.Client, error)
	Create(req *sq.CreateClientParams) (*sq.Client, error)
	GetByEmail(email string) (*sq.Client, error)
	UpdateById(arg *sq.UpdateClientByIdParams) error
}
