package repository

import (
	"context"
	"database/sql"

	gen "github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository/db/gen"
)

type UserRepo struct {
	q *gen.Queries
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{q: gen.New(db)}
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*gen.User, error) {
	user, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
