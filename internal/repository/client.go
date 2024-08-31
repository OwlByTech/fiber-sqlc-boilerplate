package repository

import (
	"context"
	"optitech/internal/interfaces"
	sq "optitech/internal/sqlc"
)

type query struct {
	repository *sq.Queries
}

func NewRepositoryClient(q *sq.Queries) interfaces.IClientRepository {
	return &query{
		repository: q,
	}
}

func (q *query) Get(id *int64) (*sq.Client, error) {
	ctx := context.Background()
	res, err := q.repository.GetClient(ctx, *id)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (q *query) Create(req *sq.CreateClientParams) (*sq.Client, error) {
	ctx := context.Background()
	res, err := q.repository.CreateClient(ctx, *req)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
