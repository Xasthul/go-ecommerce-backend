package repository

import (
	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type ProductRepository struct {
	q *gen.Queries
}

func NewProductRepository(q *gen.Queries) *ProductRepository {
	return &ProductRepository{q: q}
}
