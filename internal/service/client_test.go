package service

import (
	"optitech/database"
	dto "optitech/internal/dto/client"
	"optitech/internal/repository"
	"testing"

	sq "optitech/internal/sqlc"

	"github.com/stretchr/testify/assert"
)

func TestClientServices(t *testing.T) {
	db, err := database.Connect()

	assert.NotNil(t, db)
	assert.Nil(t, err)

	repository.Queries = *sq.New(db)

	var client *sq.Client

	t.Run("Create a client", func(t *testing.T) {
		req := dto.CreateClientReq{
			GivenName: "test",
			Surname:   "test",
			Email:     "test@gmail.com",
			Password:  "password",
		}
		client, err = CreateClientService(req)
		assert.NotNil(t, client)
		assert.Nil(t, err)
	})

	var getClient *dto.GetClientRes

	t.Run("Get the client created previously", func(t *testing.T) {
		assert.NotNil(t, client)

		req := dto.GetClientReq{
			Id: client.ClientID,
		}

		getClient, err = GetClientService(req)
		assert.NotNil(t, getClient)
		assert.Nil(t, err)
	})

	assert.NotNil(t, getClient)
	assert.NotNil(t, client)
	assert.Equal(t, getClient.Id, client.ClientID)
}
