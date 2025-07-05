package repository

import (
	"context"
	"time"

	gen "github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type TokenRepository struct {
	q *gen.Queries
}

func NewTokenRepository(q *gen.Queries) *TokenRepository {
	return &TokenRepository{q: q}
}

func (r *TokenRepository) CreateRefreshToken(ctx context.Context, userId uuid.UUID, tokenHash string, expiresAt time.Time) error {
	return r.q.CreateRefreshToken(ctx, gen.CreateRefreshTokenParams{
		UserID:    userId,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
	})
}
