package service

import (
	dto "optitech/internal/dto/client"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type service struct {
	repository interfaces.IClientRepository
}

func NewServiceClient(r interfaces.IClientRepository) interfaces.IClientService {
	return &service {
		repository: r,
	}
}

func (s *service) Get(req *dto.GetClientReq) (*dto.GetClientRes, error) {
	repoRes, err := s.repository.Get(&req.Id)

	if err != nil {
		return nil, err
	}

	return &dto.GetClientRes{
		Id:        repoRes.ClientID,
		Email:     repoRes.Email,
		GivenName: repoRes.GivenName,
		Surname:   repoRes.Surname,
	}, nil
}

func (s *service) Create(req *dto.CreateClientReq) (*dto.GetClientRes, error) {
	repoReq := sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  req.Password,
	}

	repoRes, err := s.repository.Create(&repoReq)

	if err != nil {
		return nil, err
	}

	res := &dto.GetClientRes{
		Id: repoRes.ClientID,
		GivenName: repoRes.GivenName,
		Surname: repoRes.Surname,
		Email: repoRes.Email,
	}

	return res, nil
}
