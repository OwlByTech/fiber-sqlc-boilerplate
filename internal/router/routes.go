package router

import (
	"optitech/internal/handler"
)

func (s Server) AttachRoutes() {
	r := s.app

	r.Post("/api/client", handler.CreateClientHandler)
	r.Get("/api/client/:id", handler.GetClientHandler)
}
