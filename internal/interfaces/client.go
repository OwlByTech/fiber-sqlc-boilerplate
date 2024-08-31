package interfaces

import (
	"optitech/internal/dto/client"
	sq "optitech/internal/sqlc"

	"github.com/gofiber/fiber/v2"
)

type IClientHandler interface {
	Get(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type IClientService interface {
	Get(req *dto.GetClientReq) (*dto.GetClientRes, error)
	Create(req *dto.CreateClientReq) (*dto.GetClientRes, error)
}

type IClientRepository interface {
	Get(id *int64) (*sq.Client, error)
	Create(req *sq.CreateClientParams) (*sq.Client, error)
}
