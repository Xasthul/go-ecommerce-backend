package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type OrderRepository struct {
	q *gen.Queries
}

func NewOrderRepository(q *gen.Queries) *OrderRepository {
	return &OrderRepository{q: q}
}

func (r *OrderRepository) CreateOrder(
	ctx context.Context,
	userId uuid.UUID,
	status string,
	totalCents int,
) (*gen.Order, error) {
	order, err := r.q.CreateOrder(ctx, gen.CreateOrderParams{
		UserID:     userId,
		Status:     status,
		TotalCents: int32(totalCents),
	})
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) CreateOrderItem(
	ctx context.Context,
	orderId uuid.UUID,
	productId uuid.UUID,
	quantity int,
	priceCents int,
) error {
	return r.q.CreateOrderItem(ctx, gen.CreateOrderItemParams{
		OrderID:    orderId,
		ProductID:  productId,
		Quantity:   int32(quantity),
		PriceCents: int32(priceCents),
	})
}
