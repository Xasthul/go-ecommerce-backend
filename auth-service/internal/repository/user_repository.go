package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository/db/gen"
)

type UserRepository struct {
	q *gen.Queries
}

func NewUserRepository(q *gen.Queries) *UserRepository {
	return &UserRepository{q: q}
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*gen.User, error) {
	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
