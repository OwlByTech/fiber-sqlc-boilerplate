package service

import (
	"database/sql"
	cfg "owlbytech/internal/config"
	dto "owlbytech/internal/dto/client"
	dto_mailing "owlbytech/internal/dto/mailing"
	"owlbytech/internal/interfaces"
	"owlbytech/internal/security"
	"owlbytech/internal/service/mailing"

	sq "owlbytech/internal/sqlc"
	"time"
)

type service struct {
	repository interfaces.IClientRepository
}

func NewService(r interfaces.IClientRepository) interfaces.IClientService {
	return &service{
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
		Password:  repoRes.Password,
	}, nil
}

func (s *service) Create(req *dto.CreateClientReq) (*dto.CreateClientRes, error) {
	hash, err := security.BcryptHashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	repoReq := &sq.CreateClientParams{
		GivenName: req.GivenName,
		Surname:   req.Surname,
		Email:     req.Email,
		Password:  hash,
		CreatedAt: time.Now(),
	}

	repoRes, err := s.repository.Create(repoReq)
	if err != nil {
		return nil, err
	}

	client := &dto.ClientToken{
		Id: repoRes.ClientID,
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &dto.CreateClientRes{
		Token: token,
	}, nil
}
func (s *service) UpdateById(req *dto.UpdateClientByIdReq) (bool, error) {
	client, err := s.Get(&dto.GetClientReq{Id: req.ClientId})

	if err != nil {
		return false, err
	}

	repoReq := &sq.UpdateClientByIdParams{
		ClientID:  req.ClientId,
		Email:     client.Email,
		GivenName: client.GivenName,
		Surname:   client.Surname,
		Password:  client.Password,
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	if req.Email != "" {
		repoReq.Email = req.Email
	}

	if req.Password != "" {
		hash, err := security.BcryptHashPassword(req.Password)
		if err != nil {
			return false, err
		}
		repoReq.Password = hash
	}

	if req.GivenName != "" {
		repoReq.GivenName = req.GivenName
	}

	if req.Surname != "" {
		repoReq.Surname = req.Surname
	}

	err = s.repository.UpdateById(repoReq)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) Login(req *dto.LoginClientReq) (*dto.LoginClientRes, error) {
	res, err := s.repository.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if security.BcryptCheckPasswordHash(req.Password, res.Password) != nil {
		return nil, err
	}

	client := &dto.ClientToken{
		Id: res.ClientID,
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecret)

	if err != nil {
		return nil, err
	}

	return &dto.LoginClientRes{
		Token: token,
	}, nil

}

func (s *service) ResetPassword(req dto.ResetPasswordReq) (bool, error) {
	res, err := s.repository.GetByEmail(req.Email)
	if err != nil {
		return false, err
	}
	client := &dto.ClientTokenResetPassword{
		Id:  res.ClientID,
		Exp: time.Now().Add(time.Hour / 2).Unix(),
	}

	token, err := security.JWTSign(client, cfg.Env.JWTSecretPassword)
	if err != nil {
		return false, err
	}

	if err := mailing.SendResetPassword(&dto_mailing.ResetPasswordMailingReq{
		Email:   res.Email,
		Subject: "Restablecer contraseña",
		Link:    cfg.Env.WebUrl + "/change-password?token=" + token,
	}); err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) ResetPasswordToken(req *dto.ResetPasswordTokenReq) (bool, error) {
	var payload dto.ClientTokenResetPassword
	err := security.JWTGetPayload(req.Token, cfg.Env.JWTSecretPassword, &payload)
	if err != nil {
		return false, err
	}
	client, err := s.Get(&dto.GetClientReq{Id: payload.Id})
	if err != nil {
		return false, err
	}
	hash, err := security.BcryptHashPassword(req.Password)
	if err != nil {
		return false, err
	}
	res, err := s.UpdateById(&dto.UpdateClientByIdReq{
		ClientId:  client.Id,
		Password:  hash,
		Email:     client.Email,
		GivenName: client.GivenName,
		Surname:   client.Surname,
	})
	if err != nil {
		return false, err
	}

	return res, nil
}
func (s *service) ValidateResetPasswordToken(req dto.ValidateResetPasswordTokenReq) (bool, error) {
	_, err := security.JWTVerify(req.Token, cfg.Env.JWTSecretPassword)
	if err != nil {
		return false, err
	}
	return true, nil
}
