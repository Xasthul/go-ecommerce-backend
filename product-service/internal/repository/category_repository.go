package repository

import (
	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type CategoryRepository struct {
	q *gen.Queries
}

func NewCategoryRepository(q *gen.Queries) *CategoryRepository {
	return &CategoryRepository{q: q}
}
