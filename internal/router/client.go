package router

import (
	"optitech/internal/handler"
	"optitech/internal/repository"
	service "optitech/internal/service/client"
)

func (s *Server) ClientRouter(){
	r := s.app

	repo := repository.NewRepositoryClient(repository.Queries)
	svc := service.NewServiceClient(repo)
	handler := handler.NewHandlerClient(svc)
	rg := r.Group("/api/client")
	rg.Post("/", handler.Create)
	rg.Get("/:id", handler.Get)
}
