package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app  *fiber.App
	Port uint16
}

func (s *Server) New() {
	s.app = fiber.New()
}

func (s *Server) ListenAndServe() error {
	s.AttachRoutes()
	err := s.app.Listen(fmt.Sprintf(":%d", s.Port))

	if err != nil {
		return err
	}

	return nil
}
