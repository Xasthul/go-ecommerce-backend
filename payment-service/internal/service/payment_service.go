package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository"
	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepository *repository.PaymentRepository
}

func NewOrderService(
	paymentRepository *repository.PaymentRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
	}
}

func (s *PaymentService) CreatePayment(
	ctx context.Context,
	userId uuid.UUID,
	orderId uuid.UUID,
	productId uuid.UUID,
	totalCents int,
) error {
	return s.paymentRepository.CreatePayment(ctx, orderId, userId, totalCents, "pending")
}
