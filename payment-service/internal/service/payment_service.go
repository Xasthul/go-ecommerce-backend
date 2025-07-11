package service

import (
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository"
)

type PaymentService struct {
	orderRepository *repository.PaymentRepository
}

func NewOrderService(
	orderRepository *repository.PaymentRepository,
) *PaymentService {
	return &PaymentService{
		orderRepository: orderRepository,
	}
}
