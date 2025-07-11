package service

import (
	"context"

	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/client"
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/rabbitmq"
	"github.com/Xasthul/go-ecommerce-backend/order-service/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepository *repository.OrderRepository
	productClient   *client.ProductClient
	publisher       *rabbitmq.Publisher
}

func NewOrderService(
	orderRepository *repository.OrderRepository,
	productClient *client.ProductClient,
	publisher *rabbitmq.Publisher,
) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productClient:   productClient,
		publisher:       publisher,
	}
}

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
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

	if product.Stock < quantity {
		return &AppError{Code: 409, Message: "Not enough product in stock"}
	}

	totalCents := product.PriceCents * quantity
	order, err := s.orderRepository.CreateOrder(ctx, userId, "pending", totalCents)
	if err != nil {
		return &AppError{Code: 500, Message: "Failed to create an order"}
	}

	err = s.orderRepository.CreateOrderItem(ctx, order.ID, productId, quantity, product.PriceCents)
	if err != nil {
		return &AppError{Code: 500, Message: "Failed to create an order item"}
	}

	err = s.publisher.PublishOrderCreated(&rabbitmq.OrderCreatedEvent{
		UserID:     userId,
		OrderID:    order.ID,
		ProductID:  productId,
		Quantity:   quantity,
		TotalCents: totalCents,
	})
	if err != nil {
		return &AppError{Code: 500, Message: "Failed to publish order created event"}
	}

	return nil
}
