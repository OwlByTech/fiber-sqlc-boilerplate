package repository

import (
	"context"
	"owlbytech/internal/interfaces"
	sq "owlbytech/internal/sqlc"
)

type query struct {
	repository *sq.Queries
}

func NewRepository(q *sq.Queries) interfaces.IClientRepository {
	return &query{
		repository: q,
	}
}

func (q *query) Create(req *sq.CreateClientParams) (*sq.Client, error) {
	ctx := context.Background()
	res, err := q.repository.CreateClient(ctx, *req)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (q *query) Get(id *int64) (*sq.Client, error) {
	ctx := context.Background()
	res, err := q.repository.GetClient(ctx, *id)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *query) GetByEmail(email string) (*sq.Client, error) {
	ctx := context.Background()
	res, err := r.repository.GetClientByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *query) UpdateById(arg *sq.UpdateClientByIdParams) error {
	ctx := context.Background()
	return r.repository.UpdateClientById(ctx, *arg)
}
