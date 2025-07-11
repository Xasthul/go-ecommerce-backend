package repository

import (
	"context"

	gen "github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository/db/gen"
	"github.com/google/uuid"
)

type PaymentRepository struct {
	q *gen.Queries
}

func NewPaymentRepository(q *gen.Queries) *PaymentRepository {
	return &PaymentRepository{q: q}
}

func (r *PaymentRepository) CreatePayment(
	ctx context.Context,
	orderId uuid.UUID,
	userId uuid.UUID,
	amountCents int,
	status string,
) (*gen.Payment, error) {
	payment, err := r.q.CreatePayment(ctx, gen.CreatePaymentParams{
		OrderID:     orderId,
		UserID:      userId,
		AmountCents: int32(amountCents),
		Status:      status,
	})
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
