package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/rabbitmq"
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository"
	"github.com/google/uuid"
)

type PaymentService struct {
	paymentRepository *repository.PaymentRepository
	publisher         *rabbitmq.Publisher
}

func NewOrderService(
	paymentRepository *repository.PaymentRepository,
	publisher *rabbitmq.Publisher,
) *PaymentService {
	return &PaymentService{
		paymentRepository: paymentRepository,
		publisher:         publisher,
	}
}

func (s *PaymentService) CreatePayment(
	ctx context.Context,
	userId uuid.UUID,
	orderId uuid.UUID,
	productId uuid.UUID,
	totalCents int,
) error {
	payment, err := s.paymentRepository.CreatePayment(ctx, orderId, userId, totalCents, "pending")
	if err != nil {
		s.publisher.PublishPaymentFailed(&rabbitmq.PaymentFailedEvent{
			OrderID: orderId,
			Reason:  err.Error(),
		})
		return err
	}

	s.publisher.PublishPaymentSucceeded(&rabbitmq.PaymentSucceededEvent{
		OrderID:   orderId,
		PaymentID: payment.ID,
		Amount:    totalCents,
	})
	return nil
}
