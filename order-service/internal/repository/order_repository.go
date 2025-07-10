package repository

import (
	gen "github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository/db/gen"
)

type OrderRepository struct {
	q *gen.Queries
}

func NewOrderRepository(q *gen.Queries) *OrderRepository {
	return &OrderRepository{q: q}
}
