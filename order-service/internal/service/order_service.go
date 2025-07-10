package service

import (
	"context"
	"fmt"

	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/client"
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
	productClient   *client.ProductClient
}

func NewOrderService(
	orderRepository *repository.OrderRepository,
	productClient *client.ProductClient,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productClient:   productClient,
	}
}

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	userId uuid.UUID,
	productId uuid.UUID,
	quantity int,
) error {
	product, err := s.productClient.GetProduct(ctx, productId.String())
	if err != nil {
		return &AppError{Code: 500, Message: "Failed to get a product"}
	}

	totalCents := product.PriceCents * quantity
	order, err := s.orderRepository.CreateOrder(ctx, userId, "pending", totalCents)
	if err != nil {
		return &AppError{Code: 500, Message: "Failed to create an order"}
	}

	return s.orderRepository.CreateOrderItem(ctx, order.ID, productId, quantity, product.PriceCents)
}
