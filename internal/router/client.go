package router

import (
	"owlbytech/internal/handler"
	"owlbytech/internal/middleware"
	"owlbytech/internal/repository"
	service "owlbytech/internal/service/client"
)

func (s *Server) ClientRouter() {
	r := s.app

	repo := repository.NewRepository(repository.Queries)
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	middleware := middleware.ClientMiddleware{
		Service: svc,
	}

	rg := r.Group("/api/client")
	rg.Post("/", handler.Create)
	rg.Get("/", middleware.ClientJWT, handler.Get)
	rg.Post("/login", handler.Login)
	rg.Post("/reset-password", middleware.ClientJWT, handler.ResetPassword)
	rg.Post("/reset-password-token", middleware.ClientJWT, handler.ResetPasswordToken)
	rg.Get("/validate/reset-password-token", middleware.ClientJWT, handler.ValidateResetPasswordToken)
}
