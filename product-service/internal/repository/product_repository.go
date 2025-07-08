package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
)

type ProductRepository struct {
	q *gen.Queries
}

func NewProductRepository(q *gen.Queries) *ProductRepository {
	return &ProductRepository{q: q}
}

func (r *ProductRepository) GetProducts(ctx context.Context) ([]gen.Product, error) {
	products, err := r.q.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}
