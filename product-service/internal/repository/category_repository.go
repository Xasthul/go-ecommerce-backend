package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type CategoryRepository struct {
	q *gen.Queries
}

func NewCategoryRepository(q *gen.Queries) *CategoryRepository {
	return &CategoryRepository{q: q}
}

func (r *CategoryRepository) CreateCategory(ctx context.Context, name string) (*gen.Category, error) {
	category, err := r.q.CreateCategory(ctx, name)
	if err != nil {
		return nil, err
	}
	return &category, nil
}
