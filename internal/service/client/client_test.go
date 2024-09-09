package service

import (
	"owlbytech/database"
	cfg "owlbytech/internal/config"
	dto "owlbytech/internal/dto/client"
	"owlbytech/internal/repository"
	"owlbytech/internal/security"
	"testing"

	sq "owlbytech/internal/sqlc"

	"github.com/stretchr/testify/assert"
)

func TestClientServices(t *testing.T) {
	err := cfg.LoadConfig()
	assert.Nil(t, err)

	db, err := database.Connect()

	assert.NotNil(t, db)
	assert.Nil(t, err)

	repository.Queries = sq.New(db)
	repo := repository.NewRepository(repository.Queries)
	svc := NewService(repo)
	var client *dto.CreateClientRes

	t.Run("Create a client", func(t *testing.T) {
		req := &dto.CreateClientReq{
			Email:     "test@gmail.com",
			GivenName: "test",
			Surname:   "test",
			Password:  "password",
		}
		client, err = svc.Create(req)
		assert.NotNil(t, client)
		assert.Nil(t, err)
	})

	var getClient *dto.GetClientRes
	var clientVerified dto.ClientToken

	t.Run("Get the client created previously", func(t *testing.T) {
		assert.NotNil(t, client)

		err := security.JWTGetPayload(client.Token, cfg.Env.JWTSecret, &clientVerified)

		assert.Nil(t, err)

		req := &dto.GetClientReq{
			Id: clientVerified.Id,
		}

		getClient, err = svc.Get(req)
		assert.NotNil(t, getClient)
		assert.Nil(t, err)
	})

	assert.NotNil(t, getClient)
	assert.NotNil(t, client)
	assert.Equal(t, getClient.Id, clientVerified.Id)
}
