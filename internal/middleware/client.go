package middleware

import (
	cfg "owlbytech/internal/config"
	cdto "owlbytech/internal/dto/client"
	"owlbytech/internal/interfaces"
	"owlbytech/internal/security"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ClientMiddleware struct {
	Service interfaces.IClientService
}

func (cm ClientMiddleware) ClientJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing Authorization header",
		})
	}

	splitToken := strings.Split(authHeader, " ")
	if len(splitToken) != 2 || splitToken[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := splitToken[1]
	var clientVerified cdto.ClientToken
	err := security.JWTGetPayload(token, cfg.Env.JWTSecret, &clientVerified)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "StatusUnauthorized"})
	}

	_, err = cm.Service.Get(&cdto.GetClientReq{Id: clientVerified.Id})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "StatusUnauthorized"})
	}

	c.Locals("clientId", clientVerified.Id)

	return c.Next()
}
